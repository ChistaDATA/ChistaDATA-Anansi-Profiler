package data_handlers

import "bufio"

type ILineParsingFunc interface {
	GetFunc() bufio.SplitFunc
	IsUsable(databaseType string, version string) bool
}

type LineParsingFunc struct {
	databaseType string
	minVersion   string
	maxVersion   string
	splitFunc    bufio.SplitFunc
}

func InitLineParsingFunc(databaseType string, minVersion string, maxVersion string, splitFunc bufio.SplitFunc) *LineParsingFunc {
	return &LineParsingFunc{
		databaseType: databaseType,
		minVersion:   minVersion,
		maxVersion:   maxVersion,
		splitFunc:    splitFunc,
	}
}

func (l *LineParsingFunc) GetFunc() bufio.SplitFunc {
	return l.splitFunc
}

func (l *LineParsingFunc) IsUsable(databaseType string, version string) bool {
	if databaseType == l.databaseType && version >= l.minVersion && version <= l.maxVersion {
		return true
	}
	return false
}

var lineParsingFunctions = [...]ILineParsingFunc{}

func GetSplitFunc(name string, version string) ILineParsingFunc {
	for _, function := range lineParsingFunctions {
		if function.IsUsable(name, version) {
			return function
		}
	}
	return nil
}
