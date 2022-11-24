package data_handlers

import (
	"bufio"
	log "github.com/sirupsen/logrus"
	"os"
)

type FileHandler struct {
	file      *os.File
	scanner   *bufio.Scanner
	splitFunc ILineParsingFunc
}

func InitFileHandler(filePath string, databaseType string, databaseVersion string) (*FileHandler, error) {
	fh := FileHandler{}
	var err error
	fh.file, err = os.Open(filePath)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	fh.scanner = bufio.NewScanner(fh.file)
	fh.splitFunc = GetSplitFunc(databaseType, databaseVersion)

	if fh.splitFunc != nil {
		fh.scanner.Split(fh.splitFunc.GetFunc())
	}

	return &fh, nil
}

func (fh *FileHandler) Scan() bool {
	return fh.scanner.Scan()
}

func (fh *FileHandler) Text() string {
	return fh.scanner.Text()
}

func (fh *FileHandler) Err() error {
	return fh.scanner.Err()
}

func (fh *FileHandler) Close() error {
	return fh.file.Close()
}
