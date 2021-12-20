package references

import (
	"create_jvm/main02/instructions/base"
	"create_jvm/main02/rtda"
)

type INVOKE_SPECIAL struct {
	base.Index16Instruction
} // hack!
func (self *INVOKE_SPECIAL) Execute(frame *rtda.Frame) {
	frame.OperandStack().PopRef()
}
