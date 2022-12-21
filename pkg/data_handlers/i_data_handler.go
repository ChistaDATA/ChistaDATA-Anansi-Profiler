package data_handlers

// IDataHandler is the wrapper for multiple kinds of data, local files, S3 files, remote fetch
// TODO rename to cursor
type IDataHandler interface {
	IsNextLine() bool
	GetLine() string
	Close()
}
