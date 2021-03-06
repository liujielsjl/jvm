package base

import (
	"jvm/pkg/rtda"
)

type Instruction interface {
	// 从 reader 里读取操作数，不同的指令操作数个数、类型不同
	FetchOperands(reader *ByteCodeReader)

	// 执行指令
	Execute(frame *rtda.Frame)
}

// 无操作数的指令
type NoOperandsInstruction struct {
}

func (this *NoOperandsInstruction) FetchOperands(reader *ByteCodeReader) {
	// 本身没有操作数，do nothing
}

// 操作数是单字节下标，通常是从本地变量表的下标
type Index8Instruction struct {
	Index uint
}

func (this *Index8Instruction) FetchOperands(reader *ByteCodeReader) {
	this.Index = uint(reader.ReadUint8())
}

// 操作数是两个字节的下标，通常是运行时常量表的下标
type Index16Instruction struct {
	Index uint
}

func (this *Index16Instruction) FetchOperands(reader *ByteCodeReader) {
	this.Index = uint(reader.ReadUint16())

}

type BranchInstruction struct {
	Offset int
}

func (this *BranchInstruction) FetchOperands(reader *ByteCodeReader) {
	this.Offset = int(int16(reader.ReadUint16()))
}

func BranchJump(frame *rtda.Frame, offset int) {
	pc := frame.Thread().PC()
	frame.SetNextPC(pc + offset)
}
