package data_handlers

// TODO rename to curser
type IDataHandler interface {
	IsNextLine() bool
	GetLine() string
	Close()
}
