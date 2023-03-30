package compression

import (
	"archive/tar"
	"bufio"
	"compress/gzip"
	"fmt"
	"github.com/ChistaDATA/ChistaDATA-Profiler-for-ClickHouse.git/pkg/stucts"
	"os"
	"strings"
)

type GZipCompressionHandler struct {
}

func (z GZipCompressionHandler) CanUncompress(filepath string) bool {
	return strings.HasSuffix(filepath, ".gz") || strings.HasSuffix(filepath, ".tgz")
}

func (z GZipCompressionHandler) Uncompress(filePath string) []string {
	dst, err := os.MkdirTemp(stucts.TempFolder, "*")
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

	if strings.HasSuffix(".tgz", filePath) {
		return tarExpander(tar.NewReader(gzipReader), dst)
	}

	return createUncompressedFile(filePath, dst, bufio.NewReader(gzipReader), ".gz")
}
