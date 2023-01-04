package data_handlers

import "github.com/ChistaDATA/ChistaDATA-Profiler-for-ClickHouse.git/pkg/stucts"

// GetDataHandler using configuration it selects the best possible IDataHandler
func GetDataHandler(config *stucts.Config) IDataHandler {
	if len(config.FilePaths) > 0 {
		return InitLocalFilesDataHandler(config.FilePaths, config.DatabaseName, config.DatabaseVersion)
	} else if len(config.S3Config.AccessKeyID) > 0 {
		return InitS3DataHandler(config.S3Config, config.DatabaseName, config.DatabaseVersion)
	}
	return nil
}
