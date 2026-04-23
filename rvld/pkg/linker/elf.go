package linker

import "unsafe"

// Ehdr{}：创建一个 Ehdr 类型的零值实例（所有字段初始化为零值），用于计算该类型的内存大小。
const EhdrSize = int(unsafe.Sizeof(Ehdr{}))
const ShdrSize = int(unsafe.Sizeof(Shdr{}))

type Ehdr struct {
	// 第一个是数组，无符号8位的数组
	Ident     [16]uint8
	Type      uint16
	Machine   uint16
	Version   uint32
	Entry     uint64
	PhOff     uint64
	ShOff     uint64
	Flags     uint32
	EhSize    uint16
	PhEntSize uint16
	PhNum     uint16
	ShEntSize uint16
	ShNum     uint16
	ShStrndx  uint16
}

type Shdr struct {
	Name      uint32
	Type      uint32
	Flags     uint64
	Addr      uint64
	Offset    uint64
	Size      uint64
	Link      uint32
	Info      uint32
	AddrAlign uint64
	EntSize   uint64
}
