package parsers

import (
	"github.com/ChistaDATA/ChistaDATA-Profiler-for-ClickHouse.git/pkg/stucts"
)

type ParseLogFunc func(logLine string) (stucts.ExtractedLog, error)
type InfoParserFunc func(extractedLog stucts.ExtractedLog, dBPerfInfoRepository *stucts.DBPerfInfoRepository) error

// Parser is a IParser implementation
type Parser struct {
	minVersion              string
	maxVersion              string
	database                string
	parseLog                ParseLogFunc
	infoExecuteAllFunctions []InfoParserFunc
	infoExecuteOneFunctions []InfoParserFunc
	dBPerfInfoRepository    *stucts.DBPerfInfoRepository
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

func (p *Parser) SetDBPerfInfoRepository(DBPerfInfoRepository *stucts.DBPerfInfoRepository) {
	p.dBPerfInfoRepository = DBPerfInfoRepository
}

func (p *Parser) GetDBPerfInfoRepository() *stucts.DBPerfInfoRepository {
	return p.dBPerfInfoRepository
}

func (p *Parser) Parse(logLine string) {
	log, err := p.parseLog(logLine)
	if err == nil {
		p.extractInfoFromLog(log, p.dBPerfInfoRepository)
	}
}

func (p *Parser) IsUsable(version string, database string) bool {
	if database == p.database && version >= p.minVersion && version <= p.maxVersion {
		return true
	}
	return false
}

func (p *Parser) extractInfoFromLog(extractedLog stucts.ExtractedLog, DBPerfInfoRepository *stucts.DBPerfInfoRepository) error {
	for _, function := range p.infoExecuteAllFunctions {
		function(extractedLog, DBPerfInfoRepository)
	}

	for _, function := range p.infoExecuteOneFunctions {
		err := function(extractedLog, DBPerfInfoRepository)
		if err == nil {
			return err
		}
	}

	return nil
}
