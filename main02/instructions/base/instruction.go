package base

import "create_jvm/main02/rtda"

type Instruction interface {
	//取操作数
	FetchOperands(reader *BytecodeReader)
	//执行
	Execute(frame *rtda.Frame)
}

//空指令，不执行任何操作
type Nop struct {
}

func (this *Nop) FetchOperands(reader *BytecodeReader) {
}
func (this *Nop) Execute(frame *rtda.Frame) {

}

/**
这是无操作数指令的‘父类’
*/
type NoOperandsInstruction struct {
}

func (this *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {

}

//跳转指令的‘父类‘
type BranchInstruction struct {
	Offset int
}

func (this *BranchInstruction) FetchOperands(reader *BytecodeReader) {

	this.Offset = int(reader.ReadInt16())
}

/**
这是操作数为本地变量表索引的指令的‘父类’
*/
type Index8Instruction struct {
	Index uint
}

func (this *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	this.Index = uint(reader.ReadInt8())
}

/**
这是从运行时常量池中取操作数的指令的‘父类’
*/
type Index16Instruction struct {
	Index uint
}

func (this *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	this.Index = uint(reader.ReadInt16())
}

//对接口的部分实现，就有点像抽象类了
