package compression

import (
	"bufio"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

var compressionHandlers = []CompressionHandler{
	ZipCompressionHandler{},
	GZipCompressionHandler{},
	BZip2CompressionHandler{},
	TarCompressionHandler{},
	ZXCompressionHandler{},
}

// GetUncompressedFiles gets suitable compression handler
func GetUncompressedFiles(filePath string) []string {
	for _, handler := range compressionHandlers {
		if handler.CanUncompress(filePath) == true {
			return handler.Uncompress(filePath)
		}
	}
	panic("No Compression handler Found")
}

func IsCompressed(filePath string) bool {
	for _, handler := range compressionHandlers {
		if handler.CanUncompress(filePath) == true {
			return true
		}
	}
	return false
}

type CompressionHandler interface {
	CanUncompress(filepath string) bool
	Uncompress(filepath string) []string
}

func createUncompressedFile(filePath string, dst string, d *bufio.Reader, suffix string) []string {
	_, dstFileName := filepath.Split(filePath)
	dstfilePath := filepath.Join(dst, strings.TrimSuffix(dstFileName, suffix))
	dstFile, err := os.OpenFile(dstfilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, fs.ModePerm)
	if err != nil {
		panic(err)
	}
	defer dstFile.Close()
	dstFileBW := bufio.NewWriter(dstFile)

	_, err = io.Copy(dstFileBW, d)
	if err != nil {
		panic(err)
	}
	return []string{dstfilePath}
}
