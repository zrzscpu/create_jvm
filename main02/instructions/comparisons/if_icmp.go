package comparisons

import (
	"create_jvm/main02/instructions/base"
	"create_jvm/main02/rtda"
)

//从操作数栈中取两数进行比较

type IF_ICMPEQ struct {
	base.BranchInstruction
}

func (this *IF_ICMPEQ) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopInt()
	if val1 == val2 {
		base.Branch(frame, this.Offset)
	}
}

type IF_ICMPNE struct {
	base.BranchInstruction
}

func (this *IF_ICMPNE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopInt()
	if val1 != val2 {
		base.Branch(frame, this.Offset)
	}
}

type IF_ICMPLT struct {
	base.BranchInstruction
}

func (this *IF_ICMPLT) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopInt()
	if val1 < val2 {
		base.Branch(frame, this.Offset)
	}
}

type IF_ICMPLE struct {
	base.BranchInstruction
}

func (this *IF_ICMPLE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopInt()
	if val1 <= val2 {
		base.Branch(frame, this.Offset)
	}
}

type IF_ICMPGT struct {
	base.BranchInstruction
}

func (this *IF_ICMPGT) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopInt()
	if val1 > val2 {
		base.Branch(frame, this.Offset)
	}
}

type IF_ICMPGE struct {
	base.BranchInstruction
}

func (this *IF_ICMPGE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopInt()
	if val1 >= val2 {
		base.Branch(frame, this.Offset)
	}
}
