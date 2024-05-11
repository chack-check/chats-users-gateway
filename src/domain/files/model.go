package files

type SavedFile struct {
	originalUrl       string
	originalFilename  string
	convertedUrl      *string
	convertedFilename *string
}

func (f *SavedFile) GetOriginalUrl() string {
	return f.originalUrl
}

func (f *SavedFile) GetOriginalFilename() string {
	return f.originalFilename
}

func (f *SavedFile) GetConvertedUrl() *string {
	return f.convertedUrl
}

func (f *SavedFile) GetConvertedFilename() *string {
	return f.convertedFilename
}

func NewSavedFile(originalUrl, originalFilename string, convertedUrl, convertedFilename *string) SavedFile {
	return SavedFile{
		originalUrl:       originalUrl,
		originalFilename:  originalFilename,
		convertedUrl:      convertedUrl,
		convertedFilename: convertedFilename,
	}
}
