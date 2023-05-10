package services

import (
	"github.com/ChistaDATA/ChistaDATA-Profiler-for-ClickHouse.git/pkg/data_handlers"
	"github.com/ChistaDATA/ChistaDATA-Profiler-for-ClickHouse.git/pkg/parsers"
	"github.com/ChistaDATA/ChistaDATA-Profiler-for-ClickHouse.git/pkg/stucts"
	log "github.com/sirupsen/logrus"
	"runtime"
	"sync"
)

var maxGoroutines = runtime.NumCPU()
var guard = make(chan struct{}, maxGoroutines)

// DBPerfInfoRepositoryGenerator generates *stucts.DBPerfInfoRepository
type DBPerfInfoRepositoryGenerator struct {
	dBPerfInfoRepositorys []*stucts.DBPerfInfoRepository
	parsers               []parsers.IParser
	config                *stucts.Config
	dataHandlers          []data_handlers.IDataHandler
}

func InitDBPerfInfoRepositoryGenerator(config *stucts.Config) *DBPerfInfoRepositoryGenerator {
	dataHandlers := data_handlers.GetDataHandler(config)
	var dBPerfInfoRepositorys []*stucts.DBPerfInfoRepository
	var dbParsers []parsers.IParser
	for i := 0; i < len(dataHandlers); i++ {
		dBPerfInfoRepository := stucts.InitDBPerfInfoRepository()
		dBPerfInfoRepositorys = append(dBPerfInfoRepositorys, dBPerfInfoRepository)
		parser, err := parsers.GetParser(config.DatabaseVersion, config.DatabaseName, dBPerfInfoRepository, config)
		if err != nil {
			log.Error(err)
			panic(err)
		}
		dbParsers = append(dbParsers, parser)
	}

	return &DBPerfInfoRepositoryGenerator{
		dBPerfInfoRepositorys: dBPerfInfoRepositorys,
		parsers:               dbParsers,
		config:                config,
		dataHandlers:          dataHandlers,
	}
}

func (g *DBPerfInfoRepositoryGenerator) GenerateDBPerfInfoRepository() *stucts.DBPerfInfoRepository {
	wg := sync.WaitGroup{}
	for i := 0; i < len(g.dataHandlers); i++ {
		wg.Add(1)
		guard <- struct{}{}
		lhandler := g.dataHandlers[i]
		lparser := g.parsers[i]
		go func() {
			if g.config.DatabaseName == "clickhouse" {
				g.iterateDataAndParse(lhandler, lparser)
			}
			if g.config.DatabaseName == "postgres" {
				g.iterateDataAndParsePostgres(lhandler, lparser)
			}
			wg.Done()
			<-guard
		}()
	}

	wg.Wait()

	return g.getDBPerfInfoRepository()
}

func (g *DBPerfInfoRepositoryGenerator) getDBPerfInfoRepository() *stucts.DBPerfInfoRepository {
	return stucts.CombineDBPerfInfoRepository(g.dBPerfInfoRepositorys...)
}

// iterateDataAndParse using data_handlers.IDataHandler a line is selected and processed using parsers.IParser
// parallelization is build into this using a goroutine, a guard for the number of goroutines, and number of lines to be processed
// TODO externalise parallelization parameters
func (g *DBPerfInfoRepositoryGenerator) iterateDataAndParse(dataHandler data_handlers.IDataHandler, parser parsers.IParser) {
	defer dataHandler.Close()

	var wg sync.WaitGroup
	maxLogLines := 1000
	var lines []string
	for dataHandler.IsNextLine() {
		lines = append(lines, dataHandler.GetLine())
		if len(lines) == maxLogLines {
			guard <- struct{}{}
			wg.Add(1)
			go processLines(lines, &wg, parser, &guard)
			lines = []string{}
		}
	}
	wg.Add(1)
	go processLines(lines, &wg, parser, &guard)
	wg.Wait()
}

func processLines(lines []string, wg *sync.WaitGroup, parser parsers.IParser, guard *chan struct{}) {
	for _, line := range lines {
		parser.Parse(line)
	}
	wg.Done()
	<-*guard
}

// iterateDataAndParsePostgres using data_handlers.IDataHandler a line is selected and processed using parsers.IParser
// parallelization is build into this using a goroutine, a guard for the number of goroutines, and number of lines to be processed
func (g *DBPerfInfoRepositoryGenerator) iterateDataAndParsePostgres(dataHandler data_handlers.IDataHandler, parser parsers.IParser) {
	defer dataHandler.Close()
	maxLogLines := 1000
	var lines []string
	for dataHandler.IsNextLine() {
		lines = append(lines, dataHandler.GetLine())
		if len(lines) == maxLogLines {
			processLinesPostgres(lines, parser)
			lines = []string{}
		}
	}
	processLinesPostgres(lines, parser)
}

func processLinesPostgres(lines []string, parser parsers.IParser) {
	for _, line := range lines {
		parser.Parse(line)
	}
}
