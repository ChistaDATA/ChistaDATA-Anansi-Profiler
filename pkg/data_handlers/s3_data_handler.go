package data_handlers

import (
	"github.com/ChistaDATA/ChistaDATA-Profiler-for-ClickHouse.git/pkg/stucts"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"strings"
)

// S3DataHandler It implements IDataHandler interface, acts as a wrapper for S3 files
type S3DataHandler struct {
	databaseType    string
	databaseVersion string
	s3filePaths     []string
	s3Handler       *S3Handler
	currentFilePos  int
	fileHandler     IFileHandler
}

func InitS3DataHandler(s3Config *stucts.S3Config, fileLocations []string, databaseType string, databaseVersion string) IDataHandler {

	s3Handler := InitS3Handler(s3Config)

	return &S3DataHandler{
		s3Handler:       s3Handler,
		s3filePaths:     fileLocations,
		databaseType:    databaseType,
		databaseVersion: databaseVersion,
	}
}

func (sdh *S3DataHandler) IsNextLine() bool {
	isNextLine := sdh.fileHandler != nil && sdh.fileHandler.Scan()
	if !isNextLine && sdh.isNextFile() {
		sdh.SetNewFileHandler()
		return sdh.IsNextLine()
	}
	return isNextLine
}

func (sdh *S3DataHandler) isNextFile() bool {
	return sdh.currentFilePos < len(sdh.s3filePaths)
}

func (sdh *S3DataHandler) GetLine() string {
	return sdh.fileHandler.Text()
}

func (sdh *S3DataHandler) SetNewFileHandler() {
	sdh.Close()

	//TODO what if http?
	if !strings.HasPrefix(sdh.s3filePaths[sdh.currentFilePos], "https://") {
		log.Errorln("Invalid S3 Object URL: ", sdh.s3filePaths[sdh.currentFilePos])
		sdh.currentFilePos += 1
		return
	}

	tempFile := getS3TempFileLocationFromObjectURL(sdh.s3filePaths[sdh.currentFilePos])

	dirToCreate, _ := filepath.Split(tempFile)

	err := os.MkdirAll(dirToCreate, os.ModePerm)

	if err != nil {
		log.Error(err)
		sdh.currentFilePos += 1
		return
	}

	file, err := sdh.s3Handler.DownloadToFileBig(sdh.s3filePaths[sdh.currentFilePos], tempFile)
	if err != nil {
		sdh.currentFilePos += 1
		return
	}

	sdh.fileHandler, err = InitFileHandlerWithCompression(file.Name(), sdh.databaseType, sdh.databaseVersion)
	sdh.currentFilePos += 1
	if err != nil {
		log.Error(err)
	}
}

func (sdh *S3DataHandler) Close() {
	if sdh.fileHandler != nil {
		err := sdh.fileHandler.Close()
		if err != nil {
			log.Error(err)
		}
		err = os.RemoveAll(sdh.fileHandler.GetPath())
		if err != nil {
			log.Error(err)
		}
		sdh.fileHandler = nil
	}
}

func getS3TempFileLocationFromObjectURL(s3ObjectURL string) string {
	tempFile, _ := filepath.Abs(filepath.Join(stucts.TempFolder, s3ObjectURL[len("https://"):]))
	return tempFile
}
