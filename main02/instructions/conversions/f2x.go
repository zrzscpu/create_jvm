package conversions

import (
	"create_jvm/main02/instructions/base"
	"create_jvm/main02/rtda"
)

type F2D struct {
	base.NoOperandsInstruction
}

func (this *F2D) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	double := stack.PopFloat()
	res := float64(double)
	stack.PushDouble(res)
}

type F2I struct {
	base.NoOperandsInstruction
}

func (this *F2I) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	double := stack.PopFloat()
	res := int32(double)
	stack.PushInt(res)
}

type F2L struct {
	base.NoOperandsInstruction
}

func (this *F2L) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	double := stack.PopFloat()
	res := int64(double)
	stack.PushLong(res)
}
