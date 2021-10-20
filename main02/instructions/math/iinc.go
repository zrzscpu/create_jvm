package math

import (
	"create_jvm/main02/instructions/base"
	"create_jvm/main02/rtda"
)

//iinc指令给局部变量表中的int变量增加常量值
type IINNC struct {
	Index uint
	Const int32
}

func (this *IINNC) FetchOperands(reader *base.BytecodeReader) {

	/**
	局部变量表的索引正常情况下只有256
	*/
	this.Index = uint(reader.ReadUint8())
	//The const is an immediate signed byte
	this.Const = int32(reader.ReadInt8())

}
func (this *IINNC) Execute(frame *rtda.Frame) {
	localVars := frame.LocalVars()
	val := localVars.GetInt(this.Index)
	val += this.Const
	localVars.SetInt(this.Index, val)
}
