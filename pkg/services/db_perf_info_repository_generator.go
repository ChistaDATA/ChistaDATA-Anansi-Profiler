package services

import (
	"github.com/ChistaDATA/ChistaDATA-Profiler-for-ClickHouse.git/pkg/data_handlers"
	"github.com/ChistaDATA/ChistaDATA-Profiler-for-ClickHouse.git/pkg/parsers"
	"github.com/ChistaDATA/ChistaDATA-Profiler-for-ClickHouse.git/pkg/stucts"
	log "github.com/sirupsen/logrus"
	"runtime"
	"sync"
)

// DBPerfInfoRepositoryGenerator generates *stucts.DBPerfInfoRepository
type DBPerfInfoRepositoryGenerator struct {
	dBPerfInfoRepository *stucts.DBPerfInfoRepository
	parser               parsers.IParser
	config               *stucts.CliConfig
	dataHandler          data_handlers.IDataHandler
}

func InitDBPerfInfoRepositoryGenerator(config *stucts.CliConfig) *DBPerfInfoRepositoryGenerator {
	dBPerfInfoRepository := stucts.InitDBPerfInfoRepository()
	parser, err := parsers.GetParser(config.DatabaseVersion, config.DatabaseName, dBPerfInfoRepository)
	if err != nil {
		log.Error(err)
		panic(err)
	}

	return &DBPerfInfoRepositoryGenerator{
		dBPerfInfoRepository: dBPerfInfoRepository,
		parser:               parser,
		config:               config,
		dataHandler:          data_handlers.GetDataHandler(config),
	}
}

func (g *DBPerfInfoRepositoryGenerator) GenerateDBPerfInfoRepository() *stucts.DBPerfInfoRepository {
	g.iterateDataAndParse()
	return g.dBPerfInfoRepository
}

// iterateDataAndParse using data_handlers.IDataHandler a line is selected and processed using parsers.IParser
// parallelization is build into this using a goroutine, a guard for the number of goroutines, and number of lines to be processed
// TODO externalise parallelization parameters
func (g *DBPerfInfoRepositoryGenerator) iterateDataAndParse() {
	defer g.dataHandler.Close()

	var wg sync.WaitGroup
	maxLogLines := 1000
	maxGoroutines := runtime.NumCPU()
	var lines []string
	guard := make(chan struct{}, maxGoroutines)
	for g.dataHandler.IsNextLine() {
		lines = append(lines, g.dataHandler.GetLine())
		if len(lines) == maxLogLines {
			guard <- struct{}{}
			wg.Add(1)
			go processLines(lines, &wg, g.parser, &guard)
			lines = []string{}
		}
	}
	wg.Add(1)
	go processLines(lines, &wg, g.parser, &guard)
	wg.Wait()
}

func processLines(lines []string, wg *sync.WaitGroup, parser parsers.IParser, guard *chan struct{}) {
	for _, line := range lines {
		parser.Parse(line)
	}
	wg.Done()
	<-*guard
}
