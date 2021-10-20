package conversions

import (
	"create_jvm/main02/instructions/base"
	"create_jvm/main02/rtda"
)

//d2x 系列指令把double变量强制转换成其他类型
type D2F struct {
	base.NoOperandsInstruction
}

func (this *D2F) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	double := stack.PopDouble()
	res := float32(double)
	stack.PushFloat(res)
}

type D2I struct {
	base.NoOperandsInstruction
}

func (this *D2I) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	double := stack.PopDouble()
	res := int32(double)
	stack.PushInt(res)
}

type D2L struct {
	base.NoOperandsInstruction
}

func (this *D2L) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	double := stack.PopDouble()
	res := int64(double)
	stack.PushLong(res)
}
