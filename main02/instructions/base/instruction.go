package base

import "create_jvm/main02/rtda"

type Instruction interface {
	FectchOperands(reader *BytecodeReader)
	Execute(frame *rtda.Frame, frame2 rtda.Frame)
}

type NoOperandsInstruction struct {
}

func (this *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {

}

//跳转指令
type BranchInstruction struct {
	Offset int
}

func (this *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	this.Offset = int(reader.ReadInt16())
}

//从局部变量表中加载指令
type Index8Instruction struct {
	Index uint
}

func (this *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	this.Index = uint(reader.ReadInt16())
}

//从运行时常量池中加载
type Index16Instruction struct {
	Index uint
}

func (this *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	this.Index = uint(reader.ReadInt16())
}

//对接口的部分实现，就有点像抽象类了
