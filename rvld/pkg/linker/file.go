package linker

import (
	"os"
	utils_file "rvld/pkg/utils"
)

type File struct {
	Name     string
	Contents []byte
}

// 返回一个指针
func MustNewFile(name string) *File {

	contesnts, err := os.ReadFile(name)
	utils_file.MustNo(err)
	return &File{
		Name:     name,
		Contents: contesnts,
	}
}
