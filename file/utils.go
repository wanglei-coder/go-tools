package file

import "os"

// IsExist return whether the file exists
func IsExist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

// IsFile return whether is file
func IsFile(filename string) bool {
	fi, err := os.Stat(filename)
	if err != nil {
		return false
	}
	return !fi.IsDir()
}
