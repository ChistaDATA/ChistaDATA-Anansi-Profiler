package compression

var compressionHandlers = []CompressionHandler{ZipCompressionHandler{}}

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
