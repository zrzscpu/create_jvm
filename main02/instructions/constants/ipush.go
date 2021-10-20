package constants

import (
	"create_jvm/main02/instructions/base"
	"create_jvm/main02/rtda"
)

/**
bipush指令,操作数直接存在指令当中,将其压入到操作数栈
*/
type BIPUSH struct {
	val int8
}

func (this *BIPUSH) FetchOperands(reader *base.BytecodeReader) {
	this.val = reader.ReadInt8()
}
func (this *BIPUSH) Execute(frame *rtda.Frame) {
	i := int32(this.val)
	frame.OperandStack().PushInt(i)
}

type SIPUSH struct {
	val int16
}

func (this *SIPUSH) FetchOperands(reader *base.BytecodeReader) {
	this.val = reader.ReadInt16()
}
func (this *SIPUSH) Execute(frame *rtda.Frame) {
	i := int32(this.val)
	frame.OperandStack().PushInt(i)
}
