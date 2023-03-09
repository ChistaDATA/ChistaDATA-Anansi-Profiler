package compression

import (
	"bufio"
	"compress/bzip2"
	"github.com/ChistaDATA/ChistaDATA-Profiler-for-ClickHouse.git/pkg/stucts"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

type BZip2CompressionHandler struct {
}

func (z BZip2CompressionHandler) CanUncompress(filepath string) bool {
	return strings.HasSuffix(filepath, ".bz2")
}

func (z BZip2CompressionHandler) Uncompress(filePath string) []string {
	dst, err := os.MkdirTemp(stucts.TempFolder, "*")
	if err != nil {
		panic(err)
	}
	_, dstFileName := filepath.Split(filePath)
	dstfilePath := filepath.Join(dst, strings.TrimSuffix(dstFileName, ".bz2"))
	dstFile, err := os.OpenFile(dstfilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, fs.ModePerm)
	if err != nil {
		panic(err)
	}
	defer dstFile.Close()
	dstFileBW := bufio.NewWriter(dstFile)

	zipfile, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer zipfile.Close()

	zipfileBR := bufio.NewReader(zipfile)

	bzip2Reader := bzip2.NewReader(zipfileBR)

	// create a reader, using the bzip2.reader we were passed
	d := bufio.NewReader(bzip2Reader)

	_, err = io.Copy(dstFileBW, d)
	if err != nil {
		panic(err)
	}
	return []string{dstfilePath}
}
