package comparisons

import (
	"create_jvm/main02/instructions/base"
	"create_jvm/main02/rtda"
)

//栈顶的两个引用作为判断的跳转
type IF_ACMPEQ struct {
	base.BranchInstruction
}

func (this *IF_ACMPEQ) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	ref1 := stack.PopRef()
	ref2 := stack.PopRef()
	if ref2 == ref1 {
		base.Branch(frame, this.Offset)
	}
}

type IF_ACMPNE struct {
	base.BranchInstruction
}

func (this *IF_ACMPNE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	ref1 := stack.PopRef()
	ref2 := stack.PopRef()
	if ref2 != ref1 {
		base.Branch(frame, this.Offset)
	}
}
