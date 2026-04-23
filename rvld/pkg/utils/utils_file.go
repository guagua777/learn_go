package utils_file

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
	"runtime/debug"
)

func Fatal(v any) {
	fmt.Printf("rvld: \033[0;1;31mfatal:\033[0m %v\n", v)
	debug.PrintStack()
	os.Exit(1)
}

func MustNo(err error) {
	if err != nil {
		Fatal(err)
	}
}

func Assert(condition bool) {
	if !condition {
		Fatal("assert failed")
	}
}

// 命名返回值：在函数签名中声明的返回值变量，可在函数体内直接使用和修改。
// 泛型
func Read[T any](data []byte) (val T) {
	reader := bytes.NewReader(data)
	err := binary.Read(reader, binary.LittleEndian, &val)
	MustNo(err)
	return
}
