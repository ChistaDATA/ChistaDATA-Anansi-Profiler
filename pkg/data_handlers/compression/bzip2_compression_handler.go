package compression

import (
	"archive/tar"
	"bufio"
	"compress/bzip2"
	"github.com/ChistaDATA/ChistaDATA-Profiler-for-ClickHouse.git/pkg/stucts"
	"os"
	"strings"
)

type BZip2CompressionHandler struct {
}

func (z BZip2CompressionHandler) CanUncompress(filepath string) bool {
	return strings.HasSuffix(filepath, ".bz2") || strings.HasSuffix(filepath, ".tbz2")
}

func (z BZip2CompressionHandler) Uncompress(filePath string) []string {
	dst, err := os.MkdirTemp(stucts.TempFolder, "*")
	if err != nil {
		panic(err)
	}

	zipfile, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer zipfile.Close()

	zipfileBR := bufio.NewReader(zipfile)

	bzip2Reader := bzip2.NewReader(zipfileBR)

	// create a reader, using the bzip2.reader we were passed
	d := bufio.NewReader(bzip2Reader)

	if strings.HasSuffix(".tbz2", filePath) {
		return tarExpander(tar.NewReader(bzip2Reader), dst)
	}
	return createUncompressedFile(filePath, dst, d, ".bz2")
}
