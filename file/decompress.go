package file

import (
	"go-tools/slice"
	"strings"
	"sync"

	"github.com/mholt/archiver/v3"
)

const ()

var (
	once                    sync.Once
	SupportTypeMap          map[string]UnpackFile
	SupportUnpackTypeList   []string
	SupportTypeExtensionMap = map[string][]string{
		"zip":                   {".zip"},
		"x-lz4":                 {".lz4"},
		"x-lz4-compressed-tar":  {".tar.lz4"},
		"x-compressed-tar":      {".tar.gz", ".tgz"},
		"x-xz-compressed-tar":   {".tar.xz", ".txz"},
		"x-7z-compressed":       {".7z"},
		"vnd.rar":               {".rar"},
		"x-tar":                 {".tar", ".gtar", ".gem"},
		"x-bzip":                {".bz2", ".bz"},
		"x-bzip-compressed-tar": {".tar.bz2", ".tar.bz", ".tbz2", ".tbz", ".tb2"},
		"x-xz":                  {".xz"},
	}
)

func init() {
	SupportUnpackTypeList = GetSupportUnpackTypeList()
}

func GetSupportUnpackTypeList() []string {
	once.Do(func() {
		for k := range SupportTypeExtensionMap {
			SupportUnpackTypeList = append(SupportUnpackTypeList, k)
		}
	})
	return SupportUnpackTypeList
}

func getExtension(src string, extensions []string) string {
	for _, item := range extensions {
		if strings.HasSuffix(src, item) {
			return item
		}
	}
	return ""
}

// IsCompressFile return whether a file is a compressed file
func IsCompressFile(src string) bool {
	typ, _ := GetFileType(src)
	return slice.Index(SupportUnpackTypeList, typ) != -1
}

type UnpackFile interface {
	Unpack(src, dst string) error
	GetExtension(src string) string
}

func NewUnpackFile(typ string) UnpackFile {
	if k, ok := SupportTypeMap[typ]; !ok {
		return &NormalFile{}
	} else {
		return k
	}
}

type NormalFile struct{}

func (n *NormalFile) Unpack(src, dst string) error {
	return nil
}

func (n *NormalFile) GetExtension(src string) string { return "" }

type ZipFile struct{}

func (z *ZipFile) Unpack(src, dst string) error {
	return archiver.Unarchive(src, dst)
}
func (z *ZipFile) GetExtension(src string) string {
	extensions := SupportTypeExtensionMap["zip"]
	return getExtension(src, extensions)
}

type TarFile struct{}

func (t *TarFile) Unpack(src, dst string) error {
	return archiver.Unarchive(src, dst)
}
func (t *TarFile) GetExtension(src string) string { return "" }

// zip [.zip]
// x-lz4 [.lz4]
// x-lz4-compressed-tar [.tar.lz4]
// x-compressed-tar [.tar.gz .tgz]
// x-xz-compressed-tar [.tar.xz .txz]
// x-7z-compressed [.7z]
// vnd.rar [.rar]
// x-tar [.tar .gtar .gem]
// x-bzip [.bz2 .bz]
// x-bzip-compressed-tar [.tar.bz2 .tar.bz .tbz2 .tbz .tb2]
// x-xz-compressed-tar [.tar.xz .txz]
// x-xz [.xz]
