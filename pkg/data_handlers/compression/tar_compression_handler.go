package compression

import (
	"archive/tar"
	"github.com/ChistaDATA/ChistaDATA-Profiler-for-ClickHouse.git/pkg/stucts"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type TarCompressionHandler struct {
}

func (z TarCompressionHandler) CanUncompress(filepath string) bool {
	return strings.HasSuffix(filepath, ".tar")
}

func (z TarCompressionHandler) Uncompress(filePath string) []string {
	dst, err := os.MkdirTemp(stucts.TempFolder, "*")
	if err != nil {
		panic(err)
	}

	tarFile, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer tarFile.Close()

	tarReader := tar.NewReader(tarFile)

	return tarExpander(tarReader, dst)
}

func tarExpander(tarReader *tar.Reader, dst string) []string {
	var filepaths []string
	for {
		header, err := tarReader.Next()

		switch {

		// if no more files are found return
		case err == io.EOF:
			return filepaths

		// return any other error
		case err != nil:
			panic(err)

		// if the header is nil, just skip it (not sure how this happens)
		case header == nil:
			continue
		}

		// the target location where the dir/file should be created
		targetFilePath := filepath.Join(dst, header.Name)

		// the following switch could also be done using fi.Mode(), not sure if there
		// a benefit of using one vs. the other.
		// fi := header.FileInfo()

		// check the file type
		switch header.Typeflag {

		// if its a dir and it doesn't exist create it
		case tar.TypeDir:
			if _, err := os.Stat(targetFilePath); err != nil {
				if err := os.MkdirAll(targetFilePath, 0755); err != nil {
					panic(err)
				}
			}

		// if it's a file create it
		case tar.TypeReg:
			f, err := os.OpenFile(targetFilePath, os.O_CREATE|os.O_RDWR, os.FileMode(header.Mode))
			if err != nil {
				panic(err)
			}

			// copy over contents
			if _, err := io.Copy(f, tarReader); err != nil {
				panic(err)
			}

			// manually close here after each file operation; defering would cause each file close
			// to wait until all operations have completed.
			f.Close()
			filepaths = append(filepaths, targetFilePath)
		}
	}
	return filepaths
}
