package extended

import (
	"create_jvm/main02/instructions/base"
	"create_jvm/main02/rtda"
)

type IFNULL struct {
	base.BranchInstruction
}

type IFNONNULL struct {
	base.BranchInstruction
}

func (this *IFNULL) Execute(frame *rtda.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref == nil {
		base.Branch(frame, this.Offset)
	}
}

func (this *IFNONNULL) Execute(frame *rtda.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref != nil {
		base.Branch(frame, this.Offset)
	}
}
