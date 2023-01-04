package main

import (
	"github.com/ChistaDATA/ChistaDATA-Profiler-for-ClickHouse.git/pkg/services"
	"github.com/ChistaDATA/ChistaDATA-Profiler-for-ClickHouse.git/pkg/stucts"
	log "github.com/sirupsen/logrus"
	"os"
)

var config *stucts.Config

func init() {

	// configuring application on startup
	config = stucts.InitConfigFromCli(stucts.InitializeCliConfig())

	// TODO add file log supports
	log.SetOutput(os.Stdout)

	log.SetLevel(log.Level(stucts.LogLevels[config.LogLevel]))
}

func main() {
	dBPerfInfoRepositoryGenerator := services.InitDBPerfInfoRepositoryGenerator(config)
	dBPerfInfoRepository := dBPerfInfoRepositoryGenerator.GenerateDBPerfInfoRepository()
	reportGenerator := services.InitReportGenerator(config, dBPerfInfoRepository)
	reportGenerator.GenerateReport()
}
