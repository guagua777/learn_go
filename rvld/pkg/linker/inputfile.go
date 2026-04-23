package linker

import utils_file "rvld/pkg/utils"

type InputFile struct {
	File        *File
	ElfSections []Shdr
}

func NewInputFile(file *File) InputFile {
	f := InputFile{File: file}
	if len(file.Contents) < EhdrSize {
		utils_file.Fatal("file too small")
	}

	if !CheckMagic(file.Contents) {
		utils_file.Fatal("not an ELF file")
	}

	ehdr := utils_file.Read[Ehdr](file.Contents)
	// 字节切片
	contents := file.Contents[ehdr.ShOff:]
	shdr := utils_file.Read[Shdr](contents)

	numSections := int64(ehdr.ShNum)
	if numSections == 0 {
		numSections = int64(shdr.Size)
	}

	// 初始化切片
	f.ElfSections = []Shdr{shdr}
	for numSections > 1 {
		contents = contents[ShdrSize:]
		f.ElfSections = append(f.ElfSections, utils_file.Read[Shdr](contents))
		numSections--
	}

	return f
}
