package stucts

import (
	"github.com/alecthomas/kong"
	log "github.com/sirupsen/logrus"
)

// Default config values
const (
	TopQueryCountDefault     = 10
	ReportTypeText           = "text"
	ReportTypeMD             = "md"
	ReportTypeDefault        = ReportTypeText
	MinimumQueryCountDefault = 1
)

// ReportTypes List of supported report types
var ReportTypes = [...]string{ReportTypeText, ReportTypeMD}

type CliConfig struct {
	TopQueryCount     int      `short:"n" help:"Count of queries for top x table" default:"10"`
	ReportType        string   `short:"r" help:"Report type to be generated, types: md, text" default:"text"`
	FilePaths         []string `arg:"" required:"" help:"Paths of log files" type:"existingfile"`
	MinimumQueryCount int      `short:"c" help:"Minimum no of query calls needed" default:"1"`
}

func InitializeCliConfig() CliConfig {
	cliConfig := CliConfig{}
	kong.Parse(&cliConfig)
	cliConfig.validateCliConfig()
	return cliConfig
}

// validateCliConfig Validating CliConfig inputs from user
func (cliConfig *CliConfig) validateCliConfig() {

	valid := false
	for _, s := range ReportTypes {
		if s == cliConfig.ReportType {
			valid = true
			break
		}
	}
	if !valid {
		log.Warningln("Invalid Report type, Falling back to default")
		cliConfig.ReportType = ReportTypeDefault
	}

	valid = false
	if cliConfig.TopQueryCount <= 0 {
		log.Warningln("Invalid Top Query Count, Falling back to default")
		cliConfig.TopQueryCount = TopQueryCountDefault
	}

	if cliConfig.MinimumQueryCount <= 0 {
		log.Warningln("Invalid Minimum Query Count, Falling back to default")
		cliConfig.MinimumQueryCount = MinimumQueryCountDefault
	}
}
