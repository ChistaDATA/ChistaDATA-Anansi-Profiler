package parsers

import (
	"github.com/ChistaDATA/ChistaDATA-Profiler-for-ClickHouse.git/pkg/stucts"
)

type ParseLogFunc func(logLine string) (stucts.ExtractedLog, error)
type InfoParserFunc func(extractedLog stucts.ExtractedLog, infoCorpus *stucts.InfoCorpus) error

type Parser struct {
	minVersion              string
	maxVersion              string
	database                string
	parseLog                ParseLogFunc
	infoExecuteAllFunctions []InfoParserFunc
	infoExecuteOneFunctions []InfoParserFunc
	infoCorpus              *stucts.InfoCorpus
}

func InitParser(minVersion string, maxVersion string, database string, parseLog ParseLogFunc, infoExecuteAllFunctions []InfoParserFunc, infoExecuteOneFunctions []InfoParserFunc) *Parser {
	return &Parser{
		minVersion:              minVersion,
		maxVersion:              maxVersion,
		database:                database,
		parseLog:                parseLog,
		infoExecuteAllFunctions: infoExecuteAllFunctions,
		infoExecuteOneFunctions: infoExecuteOneFunctions,
	}
}

func (p *Parser) SetInfoCorpus(infoCorpus *stucts.InfoCorpus) {
	p.infoCorpus = infoCorpus
}

func (p *Parser) GetInfoCorpus() *stucts.InfoCorpus {
	return p.infoCorpus
}

func (p *Parser) Parse(logLine string) {
	log, err := p.parseLog(logLine)
	if err == nil {
		p.extractInfoFromLog(log, p.infoCorpus)
	}
}

func (p *Parser) IsUsable(version string, database string) bool {
	if database == p.database && version >= p.minVersion && version <= p.maxVersion {
		return true
	}
	return false
}

func (p *Parser) extractInfoFromLog(extractedLog stucts.ExtractedLog, infoCorpus *stucts.InfoCorpus) error {
	for _, function := range p.infoExecuteAllFunctions {
		function(extractedLog, infoCorpus)
	}

	for _, function := range p.infoExecuteOneFunctions {
		err := function(extractedLog, infoCorpus)
		if err == nil {
			return err
		}
	}

	return nil
}
