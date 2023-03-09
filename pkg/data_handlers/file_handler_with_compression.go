package data_handlers

import (
	"github.com/ChistaDATA/ChistaDATA-Profiler-for-ClickHouse.git/pkg/data_handlers/compression"
	log "github.com/sirupsen/logrus"
)

// FileHandlerWithCompression It implements IFileHandler interface, acts as a wrapper for file
type FileHandlerWithCompression struct {
	databaseType    string
	databaseVersion string
	filePaths       []string
	currentFilePos  int
	fileHandler     IFileHandler
	filePath        string
}

func (fh *FileHandlerWithCompression) GetPath() string {
	return fh.filePath
}

func (fh *FileHandlerWithCompression) Err() error {
	//TODO implement me
	panic("implement me")
}

func InitFileHandlerWithCompression(filePath string, databaseType string, databaseVersion string) (IFileHandler, error) {
	return &FileHandlerWithCompression{
		filePaths:       []string{filePath},
		databaseType:    databaseType,
		databaseVersion: databaseVersion,
		filePath:        filePath,
	}, nil
}

func (fh *FileHandlerWithCompression) Scan() bool {
	isNextLine := fh.fileHandler != nil && fh.fileHandler.Scan()
	if !isNextLine && fh.isNextFile() {
		fh.SetNewFileHandler()
		return fh.Scan()
	}
	return isNextLine
}

func (fh *FileHandlerWithCompression) isNextFile() bool {
	return fh.currentFilePos < len(fh.filePaths)
}

func (fh *FileHandlerWithCompression) Text() string {
	return fh.fileHandler.Text()
}

func (fh *FileHandlerWithCompression) SetNewFileHandler() {
	fh.Close()
	var err error
	if fh.isNextFile() {
		if compression.IsCompressed(fh.filePaths[fh.currentFilePos]) {
			fh.filePaths = append(fh.filePaths, compression.GetUncompressedFiles(fh.filePaths[fh.currentFilePos])...)
			fh.currentFilePos += 1
			fh.SetNewFileHandler()
		} else {
			fh.fileHandler, err = InitFileHandler(fh.filePaths[fh.currentFilePos], fh.databaseType, fh.databaseVersion)
			fh.currentFilePos += 1
			if err != nil {
				log.Error(err)
			}
		}
	}
}

func (fh *FileHandlerWithCompression) Close() error {
	if fh.fileHandler != nil {
		err := fh.fileHandler.Close()
		if err != nil {
			log.Error(err)
		}
	}
	return nil
}
