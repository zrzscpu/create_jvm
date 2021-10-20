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
这是不用取操作数的指令的‘父类’
*/
type NoOperandsInstruction struct {
}

func (this *NoOperandsInstruction) FetchOperands(reader *BytecodeReader) {

}

//跳转指令
type BranchInstruction struct {
	Offset1 uint8
	Offset2 uint8
	Offset  int
}

func (this *BranchInstruction) FetchOperands(reader *BytecodeReader) {
	this.Offset1 = reader.ReadUint8()
	this.Offset2 = reader.ReadUint8()
	i1 := int16(this.Offset1)
	i2 := int16(this.Offset2)
	////如果不转成i1，i2这样的类型得到的是243	1111 0011
	i := int16(i1<<8 | i2)
	//i := int16(this.Offset1 << 8 | this.Offset2)
	this.Offset = int(i)
}

/**
这是从当前字节码指令中中取操作数的局部变量表索引指令的‘父类’
*/
type Index8Instruction struct {
	Index uint
}

func (this *Index8Instruction) FetchOperands(reader *BytecodeReader) {
	this.Index = uint(reader.ReadInt8())
}

/**
这是从当前字节码指令中中取操作数在运行时常量池的索引	指令的‘父类’
*/
type Index16Instruction struct {
	Index uint
}

func (this *Index16Instruction) FetchOperands(reader *BytecodeReader) {
	this.Index = uint(reader.ReadInt16())
}

//对接口的部分实现，就有点像抽象类了
