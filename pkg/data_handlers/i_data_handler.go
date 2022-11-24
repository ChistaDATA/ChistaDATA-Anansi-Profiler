package data_handlers

type IDataHandler interface {
	IsNextLine() bool
	GetLine() string
	Close()
}
