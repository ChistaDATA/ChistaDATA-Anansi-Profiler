package compression

import (
	"archive/zip"
	"fmt"
	"github.com/ChistaDATA/ChistaDATA-Profiler-for-ClickHouse.git/pkg/stucts"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type ZipCompressionHandler struct {
}

func (z ZipCompressionHandler) CanUncompress(filepath string) bool {
	return strings.HasSuffix(filepath, ".zip")
}

func (z ZipCompressionHandler) Uncompress(filePath string) []string {
	filepaths := []string{}
	dst, err := os.MkdirTemp(stucts.TempFolder, "*")
	if err != nil {
		panic(err)
	}
	archive, err := zip.OpenReader(filePath)
	if err != nil {
		panic(err)
	}
	defer archive.Close()
	for _, f := range archive.File {
		filePath := filepath.Join(dst, f.Name)
		fmt.Println("unzipping file ", filePath)

		if !strings.HasPrefix(filePath, filepath.Clean(dst)+string(os.PathSeparator)) {
			fmt.Println("invalid file path")
			continue
		}
		if f.FileInfo().IsDir() {
			fmt.Println("creating directory...")
			err := os.MkdirAll(filePath, os.ModePerm)
			if err != nil {
				panic(err)
			}
			continue
		}

		if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
			panic(err)
		}

		dstFile, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			panic(err)
		}

		fileInArchive, err := f.Open()
		if err != nil {
			panic(err)
		}

		if _, err := io.Copy(dstFile, fileInArchive); err != nil {
			panic(err)
		}

		if err := dstFile.Close(); err != nil {
			panic(err)
		}

		if err := fileInArchive.Close(); err != nil {
			panic(err)
		}
		filepaths = append(filepaths, filePath)
	}
	return filepaths
}
