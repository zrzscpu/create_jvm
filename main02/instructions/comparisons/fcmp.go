package comparisons

import (
	"create_jvm/main02/instructions/base"
	"create_jvm/main02/rtda"
)

type FCMPG struct {
	base.NoOperandsInstruction
}

type FCMPL struct {
	base.NoOperandsInstruction
}

func _fcmp(frame *rtda.Frame, gFlag bool) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else if v1 < v2 {
		stack.PushInt(-1)
	} else if gFlag {
		stack.PushInt(1)
	} else {
		stack.PushInt(-1)
	}
}

//当两个float变量中至少有一个是NaN时，
//用fcmpg指 令比较的结果是1，
//而用fcmpl指令比较的结果是-1
func (this *FCMPL) Execute(frame *rtda.Frame) {
	_fcmp(frame, true)
}

func (this *FCMPG) Execute(frame *rtda.Frame) {
	_fcmp(frame, false)
}
