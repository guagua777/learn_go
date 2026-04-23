package main

import (
	"fmt"
	"os"

	// package 路径
	// 相当于
	// 指针 所在位置
	"rvld/pkg/linker"
	utils_file "rvld/pkg/utils"
)

func main() {
	fmt.Println("Hello, World!")
	// fmt.Println("%v", os.Args)
	fmt.Println(os.Args) // 直接打印，不需要 %v

	if len(os.Args) < 2 {
		// println("worng args")
		// os.Exit(1)
		utils_file.Fatal("worng args")
	}

	file := linker.MustNewFile(os.Args[1])

	if !linker.CheckMagic(file.Contents) {
		utils_file.Fatal("not an elf file")
	}

	inputFile := linker.NewInputFile(file)
	utils_file.Assert(len(inputFile.ElfSections) == 11)

}
