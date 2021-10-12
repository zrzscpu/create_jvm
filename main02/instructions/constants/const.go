package constants

import (
	"create_jvm/main02/instructions/base"
	"create_jvm/main02/rtda"
)

//这一系列指令将隐含在操作码中的常量值推入操作数栈
type ACONST_NULL struct {
	base.NoOperandsInstruction
}

type DCONST_0 struct {
	base.NoOperandsInstruction
}

type DCONST_1 struct {
	base.NoOperandsInstruction
}

type FCONST_0 struct {
	base.NoOperandsInstruction
}

type FCONST_1 struct {
	base.NoOperandsInstruction
}

type FCONST_2 struct {
	base.NoOperandsInstruction
}

type ICONST_M1 struct {
	base.NoOperandsInstruction
}

type ICONST_0 struct {
	base.NoOperandsInstruction
}

type ICONST_1 struct {
	base.NoOperandsInstruction
}

type ICONST_2 struct {
	base.NoOperandsInstruction
}

type ICONST_3 struct {
	base.NoOperandsInstruction
}

type ICONST_4 struct {
	base.NoOperandsInstruction
}

type ICONST_5 struct {
	base.NoOperandsInstruction
}

type LCONST_0 struct {
	base.NoOperandsInstruction
}

type LCONST_1 struct {
	base.NoOperandsInstruction
}

func (this *ACONST_NULL) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushRef(nil)
}
