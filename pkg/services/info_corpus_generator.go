package services

import (
	"github.com/ChistaDATA/ChistaDATA-Profiler-for-ClickHouse.git/pkg/data_handlers"
	"github.com/ChistaDATA/ChistaDATA-Profiler-for-ClickHouse.git/pkg/parsers"
	"github.com/ChistaDATA/ChistaDATA-Profiler-for-ClickHouse.git/pkg/stucts"
	log "github.com/sirupsen/logrus"
	"runtime"
	"sync"
)

type InfoCorpusGenerator struct {
	infoCorpus  *stucts.InfoCorpus
	parser      parsers.IParser
	config      *stucts.CliConfig
	dataHandler data_handlers.IDataHandler
}

func InitInfoCorpusGenerator(config *stucts.CliConfig) *InfoCorpusGenerator {
	infoCorpus := stucts.InitInfoCorpus()
	parser, error := parsers.GetParser(config.DatabaseVersion, config.DatabaseName)
	if error != nil {
		log.Error(error)
		panic(error)
	}
	parser.SetInfoCorpus(infoCorpus)

	return &InfoCorpusGenerator{
		infoCorpus:  infoCorpus,
		parser:      parser,
		config:      config,
		dataHandler: data_handlers.GetDataHandler(config),
	}
}

func (icg *InfoCorpusGenerator) GenerateInfoCorpus() *stucts.InfoCorpus {
	icg.iterateDataAndParse()
	return icg.infoCorpus
}

// readFileAndParseLogs Reads a file, extracts queries line by line
func (icg *InfoCorpusGenerator) iterateDataAndParse() {
	defer icg.dataHandler.Close()

	var wg sync.WaitGroup
	maxLogLines := 1000
	maxGoroutines := runtime.NumCPU()
	var lines []string
	guard := make(chan struct{}, maxGoroutines)
	for icg.dataHandler.IsNextLine() {
		lines = append(lines, icg.dataHandler.GetLine())
		if len(lines) == maxLogLines {
			guard <- struct{}{}
			wg.Add(1)
			go processLines(lines, &wg, icg.parser, &guard)
			lines = []string{}
		}
	}
	wg.Add(1)
	go processLines(lines, &wg, icg.parser, &guard)
	wg.Wait()

}

func processLines(lines []string, wg *sync.WaitGroup, parser parsers.IParser, guard *chan struct{}) {
	for _, line := range lines {
		parser.Parse(line)
	}
	wg.Done()
	<-*guard
}
