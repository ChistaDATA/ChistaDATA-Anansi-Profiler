package main

import (
	"github.com/ChistaDATA/ChistaDATA-Profiler-for-ClickHouse.git/pkg/services"
	"github.com/ChistaDATA/ChistaDATA-Profiler-for-ClickHouse.git/pkg/stucts"
	log "github.com/sirupsen/logrus"
	"os"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	//log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout) // TODO add file log supports

	// Only log the warning severity or above.
	log.SetLevel(log.ErrorLevel) // TODO make configurable
}

func main() {
	cliConfig := stucts.InitializeCliConfig()
	infoCorpusGenerator := services.InitInfoCorpusGenerator(&cliConfig)
	infoCorpus := infoCorpusGenerator.GenerateInfoCorpus()
	reportGenerator := services.InitReportGenerator(&cliConfig, infoCorpus)
	reportGenerator.GenerateReport()
}
