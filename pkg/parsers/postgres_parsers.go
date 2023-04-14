package parsers

import (
	"github.com/ChistaDATA/ChistaDATA-Profiler-for-ClickHouse.git/pkg/parsers/postgres"
	"github.com/ChistaDATA/ChistaDATA-Profiler-for-ClickHouse.git/pkg/stucts"
	"github.com/google/uuid"
	"time"
)

// InitPostgresParserV1 is Parser object which is a IParser implementation for postgres version 0 - TODO add postgreSQL version
func InitPostgresParserV1() IParser {
	return NewPostgresParser(
		"0",
		"0",
		stucts.PostgresDatabase,
		postgres.ParseLogV1,
		[]InfoParserFunc{
			postgres.ParseMessageWithNewQueryV1,
			postgres.ParseMessageWithQueryDurationV1,
			postgres.ParseLogMessageWithErrorRegExV1,
			//postgres.ParseLogMessageWithEndQueryRegExV1,
			postgres.ParseLogMessageWithSystemUsageRegExV1,
			postgres.ParseMessageDurationAndQueryV1,
		},
		[]InfoParserFunc{},
		func(i ...interface{}) error {
			config := i[0].(*stucts.Config)
			err := postgres.SetParseLogV1Params(config.LogPrefix)
			if err != nil {
				return err
			}
			return nil
		},
	)
}

type PostgresParser struct {
	*Parser
	CurrentLog *stucts.ExtractedLog
}

func NewPostgresParser(minVersion string, maxVersion string, database string, parseLog ParseLogFunc, infoExecuteAllFunctions []InfoParserFunc, infoExecuteOneFunctions []InfoParserFunc, initFunc InitParserFunc) IParser {
	return &PostgresParser{
		Parser: NewParser(
			minVersion,
			maxVersion,
			database,
			parseLog,
			infoExecuteAllFunctions,
			infoExecuteOneFunctions,
			initFunc,
		).(*Parser),
		CurrentLog: &stucts.ExtractedLog{},
	}
}

func (p *PostgresParser) Parse(logLine string) {
	log, err := p.parseLog(logLine)
	if err == nil {
		var time time.Time
		if log.LogLevel == "" && log.ProcessID == 0 && log.QueryId == "" && log.ThreadId == 0 && log.Timestamp == time {
			p.CurrentLog.Message += log.Message
		} else {
			p.extractInfoFromLog(*p.CurrentLog, p.dBPerfInfoRepository)
			p.CurrentLog = &log
			p.CurrentLog.QueryId = getNewUUID()
		}
	}
}

func getNewUUID() string {
	newUUID, _ := uuid.NewUUID()
	return newUUID.String()
}
