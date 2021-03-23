package file

import (
	"os"
	"path/filepath"

	"github.com/zRedShift/mimemagic"
)

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

func GetAllFiles(dirname string) (files []string, err error) {
	// 判断是文件还是文件夹
	fi, _ := os.Stat(dirname)
	if !fi.IsDir() {
		return []string{dirname}, err
	}
	files = make([]string, 0)
	err = filepath.Walk(dirname, func(filename string, fi os.FileInfo, err error) error { //遍历目录
		if err != nil { //忽略错误
			return err
		}
		if fi.IsDir() { // 忽略目录
			return nil
		}
		files = append(files, filename)
		return nil
	})
	return files, err
}
