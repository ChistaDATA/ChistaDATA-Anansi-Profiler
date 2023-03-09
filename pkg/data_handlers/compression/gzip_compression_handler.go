package compression

import (
	"compress/gzip"
	"fmt"
	"github.com/ChistaDATA/ChistaDATA-Profiler-for-ClickHouse.git/pkg/stucts"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type GZipCompressionHandler struct {
}

func (z GZipCompressionHandler) CanUncompress(filepath string) bool {
	return strings.HasSuffix(filepath, ".gz")
}

func (z GZipCompressionHandler) Uncompress(filePath string) []string {
	dst, err := os.MkdirTemp(stucts.TempFolder, "*")
	if err != nil {
		panic(err)
	}
	_, dstFileName := filepath.Split(filePath)
	dstfilePath := filepath.Join(dst, strings.TrimSuffix(dstFileName, ".gz"))
	dstFile, err := os.OpenFile(dstfilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, fs.ModePerm)
	if err != nil {
		panic(err)
	}

	gzipfile, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	gzipReader, err := gzip.NewReader(gzipfile)
	if err != nil {
		panic(err)
	}
	defer gzipReader.Close()
	_, err = io.Copy(dstFile, gzipReader)
	if err != nil {
		panic(err)
	}
	return []string{dstfilePath}
}
