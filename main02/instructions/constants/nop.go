package constants

import (
	"create_jvm/main02/instructions/base"
	"create_jvm/main02/rtda"
)

type Nop struct {
	base.NoOperandsInstruction
}

func (this *Nop) Execute(frame *rtda.Frame) {

}
