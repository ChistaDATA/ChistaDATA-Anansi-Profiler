package stucts

var TempFolder string

type Config struct {
	TopQueryCount         int
	ReportType            string
	FilePaths             []string
	MinimumQueryCallCount int
	DatabaseName          string
	DatabaseVersion       string
	SortField             string
	SortFieldOperation    string
	SortOrder             string
	LogLevel              string
	S3Config              *S3Config
	LogPrefix             string
}

func InitConfigFromCli(cliConfig *CliConfig) *Config {
	return &Config{
		TopQueryCount:         cliConfig.TopQueryCount,
		ReportType:            cliConfig.ReportType,
		FilePaths:             cliConfig.FilePaths,
		MinimumQueryCallCount: cliConfig.MinimumQueryCallCount,

		DatabaseName:       cliConfig.DatabaseName,
		DatabaseVersion:    cliConfig.DatabaseVersion,
		SortField:          cliConfig.SortField,
		SortFieldOperation: cliConfig.SortFieldOperation,
		SortOrder:          cliConfig.SortOrder,
		LogLevel:           cliConfig.LogLevel,
		S3Config: &S3Config{
			AccessKeyID:     cliConfig.S3AccessKeyID,
			SecretAccessKey: cliConfig.S3SecretAccessKey,
			SessionToken:    cliConfig.S3SessionToken,
			Region:          cliConfig.S3Region,
			FileLocations:   cliConfig.S3FileLocations,
		},
	}
}
