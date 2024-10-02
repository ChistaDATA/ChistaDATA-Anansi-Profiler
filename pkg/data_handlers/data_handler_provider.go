package data_handlers

import "github.com/ChistaDATA/ChistaDATA-Profiler-for-ClickHouse.git/pkg/stucts"

// GetDataHandler using configuration it selects the best possible IDataHandler
func GetDataHandler(config *stucts.Config) []IDataHandler {
	var handlers []IDataHandler
	if len(config.FilePaths) > 0 {
		for _, path := range config.FilePaths {
			handlers = append(handlers, InitLocalFilesDataHandler([]string{path}, config.DatabaseType, config.DatabaseVersion))
		}
	}
	if len(config.S3Config.AccessKeyID) > 0 {
		handlers = append(handlers, InitS3DataHandler(&config.S3Config, config.S3Config.FileLocations, config.DatabaseType, config.DatabaseVersion))
	}
	return handlers
}
