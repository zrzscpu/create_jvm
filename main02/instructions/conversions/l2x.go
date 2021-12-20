package conversions

import (
	"create_jvm/main02/instructions/base"
	"create_jvm/main02/rtda"
)

type L2F struct {
	base.NoOperandsInstruction
}

func (this *L2F) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopLong()
	stack.PushFloat(float32(val))
}

type L2D struct {
	base.NoOperandsInstruction
}

func (this *L2D) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopLong()
	stack.PushDouble(float64(val))
}

type L2I struct {
	base.NoOperandsInstruction
}

func (this *L2I) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopLong()
	stack.PushInt(int32(val))
}
