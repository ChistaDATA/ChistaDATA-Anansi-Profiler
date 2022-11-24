package data_handlers

import (
	log "github.com/sirupsen/logrus"
)

type LocalFilesDataHandler struct {
	databaseType    string
	databaseVersion string
	filePaths       []string
	currentFilePos  int
	fileHandler     *FileHandler
}

func InitLocalFilesDataHandler(filePaths []string, databaseType string, databaseVersion string) *LocalFilesDataHandler {
	return &LocalFilesDataHandler{
		filePaths:       filePaths,
		databaseType:    databaseType,
		databaseVersion: databaseVersion,
	}
}

func (fh *LocalFilesDataHandler) IsNextLine() bool {
	isNextLine := fh.fileHandler != nil && fh.fileHandler.Scan()
	if !isNextLine && fh.isNextFile() {
		fh.SetNewFileHandler()
		return fh.IsNextLine()
	}
	return isNextLine
}

func (fh *LocalFilesDataHandler) isNextFile() bool {
	return fh.currentFilePos < len(fh.filePaths)
}

func (fh *LocalFilesDataHandler) GetLine() string {
	return fh.fileHandler.Text()
}

func (fh *LocalFilesDataHandler) SetNewFileHandler() {
	fh.Close()

	var err error
	fh.fileHandler, err = InitFileHandler(fh.filePaths[fh.currentFilePos], fh.databaseType, fh.databaseVersion)
	fh.currentFilePos += 1
	if err != nil {
		log.Error(err)
	}
}

func (fh *LocalFilesDataHandler) Close() {
	if fh.fileHandler != nil {
		err := fh.fileHandler.Close()
		if err != nil {
			log.Error(err)
		}
	}
}
