package comparisons

import (
	"create_jvm/main02/instructions/base"
	"create_jvm/main02/rtda"
)

//if跳转指令
//·ifeq：x==0
//·ifne：x！=0
//·iflt：x<0
//·ifle：x<=0
//·ifgt：x>0
//·ifge：x>=0
//如果判断的过程嵌入到if里面那么if这一系列就需要 4*6中实现
//而如果拆开 只需要4+6中实现

//从操作数栈中取两数的比较的结果决定是否跳转
//如果相等
type IFEQ struct {
	base.BranchInstruction
}

func (this *IFEQ) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	if val == 0 {
		base.Branch(frame, this.Offset)
	}
}

type IFNE struct {
	base.BranchInstruction
}

func (this *IFNE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	if val != 0 {
		base.Branch(frame, this.Offset)
	}
}

type IFLT struct {
	base.BranchInstruction
}

func (this *IFLT) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	if val < 0 {
		base.Branch(frame, this.Offset)
	}
}

type IFLE struct {
	base.BranchInstruction
}

func (this *IFLE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	if val <= 0 {
		base.Branch(frame, this.Offset)
	}
}

type IFGT struct {
	base.BranchInstruction
}

func (this *IFGT) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	if val > 0 {
		base.Branch(frame, this.Offset)
	}
}

type IFGE struct {
	base.BranchInstruction
}

func (this *IFGE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val := stack.PopInt()
	if val >= 0 {
		base.Branch(frame, this.Offset)
	}
}
