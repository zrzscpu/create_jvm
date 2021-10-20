package extended

import (
	"create_jvm/main02/instructions/base"
	"create_jvm/main02/rtda"
)

//与goto的区别只是offset从2个字节变成了四个字节
type GOTO_W struct {
	offset int
}

func (this *GOTO_W) FetchOperands(reader *base.BytecodeReader) {
	this.offset = int(reader.ReadInt32())
}

func (this *GOTO_W) Execute(frame *rtda.Frame) {
	base.Branch(frame, this.offset)
}
