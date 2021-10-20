package math

import (
	"create_jvm/main02/instructions/base"
	"create_jvm/main02/rtda"
)

/**
布尔运算
*/
type IAND struct {
	base.NoOperandsInstruction
}
type LAND struct {
	base.NoOperandsInstruction
}

func (this *IAND) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopInt()
	res := val1 & val2
	stack.PushInt(res)
}

func (this *LAND) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopLong()
	val1 := stack.PopLong()
	res := val1 & val2
	stack.PushLong(res)
}
