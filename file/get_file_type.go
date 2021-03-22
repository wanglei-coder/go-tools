package file

import "github.com/zRedShift/mimemagic"

const (
	Unknow = "Unknow"
)

// GetFileType return file type
func GetFileType(filename string) (string, error) {
	mimeType, err := mimemagic.MatchFilePath(filename, -1)
	if err != nil {
		return Unknow, err
	}
	// mimeType.Subtype
	return mimeType.Subtype, nil
}
