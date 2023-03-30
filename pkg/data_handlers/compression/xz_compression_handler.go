package compression

import (
	"archive/tar"
	"bufio"
	"github.com/ChistaDATA/ChistaDATA-Profiler-for-ClickHouse.git/pkg/stucts"
	"github.com/ulikunitz/xz"
	"os"
	"strings"
)

type ZXCompressionHandler struct {
}

func (z ZXCompressionHandler) CanUncompress(filepath string) bool {
	return strings.HasSuffix(filepath, ".xz") || strings.HasSuffix(filepath, ".txz")
}

func (z ZXCompressionHandler) Uncompress(filePath string) []string {
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

	xzReader, err := xz.NewReader(zipfileBR)
	if err != nil {
		panic(err)
	}

	// create a reader, using the bzip2.reader we were passed
	xzBufferedReader := bufio.NewReader(xzReader)

	if strings.HasSuffix(".txz", filePath) {
		return tarExpander(tar.NewReader(xzReader), dst)
	}

	return createUncompressedFile(filePath, dst, xzBufferedReader, ".xz")
}
