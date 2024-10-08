package main

import (
	"github.com/ChistaDATA/ChistaDATA-Profiler-for-ClickHouse.git/pkg/services"
	"github.com/ChistaDATA/ChistaDATA-Profiler-for-ClickHouse.git/pkg/stucts"
	log "github.com/sirupsen/logrus"
	"os"
)

var config *stucts.Config

func init() {
	openTemp()
	config = stucts.InitializeCliConfig()

	// TODO add file log supports
	log.SetOutput(os.Stdout)

	level, err := log.ParseLevel(config.LogLevel)
	if err != nil {
		log.Fatalln("Invalid log level", err)
	}
	log.SetLevel(level)
}

func openTemp() {
	dirName, err := os.MkdirTemp(".", "__TempFile-*__")
	if err != nil {
		panic(err)
	}
	stucts.TempFolder = dirName
}

func closeTemp() {
	err := os.RemoveAll(stucts.TempFolder)
	if err != nil {
		panic(err)
	}
}

func main() {
	defer closeTemp()
	dBPerfInfoRepositoryGenerator := services.InitDBPerfInfoRepositoryGenerator(config)
	dBPerfInfoRepository := dBPerfInfoRepositoryGenerator.GenerateDBPerfInfoRepository()
	reportGenerator := services.InitReportGenerator(config, dBPerfInfoRepository)
	reportGenerator.GenerateReport()
}
