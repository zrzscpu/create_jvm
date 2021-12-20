package conversions

import (
	"create_jvm/main02/instructions/base"
	"create_jvm/main02/rtda"
)

type I2F struct {
	base.NoOperandsInstruction
}

func (this *I2F) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	stack.PushFloat(float32(val))
}

type I2D struct {
	base.NoOperandsInstruction
}

func (this *I2D) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	stack.PushDouble(float64(val))
}

type I2L struct {
	base.NoOperandsInstruction
}

func (this *I2L) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	stack.PushLong(int64(val))
}
