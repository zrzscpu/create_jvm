package comparisons

import (
	"create_jvm/main02/instructions/base"
	"create_jvm/main02/rtda"
)

//lcmp指令用于比较long变量
type LCMP struct {
	base.NoOperandsInstruction
}

func (this *LCMP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	val2 := stack.PopLong()
	val1 := stack.PopLong()
	//TRUE用1表示
	if val1 > val2 {
		stack.PushInt(1)
	} else if val1 == val2 {
		stack.PushInt(0)
	} else {
		stack.PushInt(-1)
	}

}
