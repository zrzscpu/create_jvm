package math

import (
	"create_jvm/main02/instructions/base"
	"create_jvm/main02/rtda"
)

type IADD struct {
	base.NoOperandsInstruction
}

func (this *IADD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopInt()
	stack.PushInt(val1 + val2)
}
