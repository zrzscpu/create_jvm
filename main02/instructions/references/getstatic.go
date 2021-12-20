package references

import (
	"create_jvm/main02/instructions/base"
	"create_jvm/main02/rtda"
	"create_jvm/main02/rtda/heap"
)

type GET_STATIC struct {
	base.Index16Instruction
}

func (this *GET_STATIC) Execute(frame *rtda.Frame) {
	pool := frame.Method().Class().ConstantPool()
	//在常量池中查找改域的符号引用
	fieldRef := pool.GetConstant(this.Index).(*heap.FieldRef)
	//查找到后找真正的域
	field := fieldRef.ResolvedField()
	class := field.Class()
	if !field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	//拿到域的描述符
	descriptor := field.Descriptor()
	slotId := field.SlotId()
	//域的所有static属性
	slots := class.StaticVars()
	//将static属性放到栈中
	stack := frame.OperandStack()
	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I':
		stack.PushInt(slots.GetInt(slotId))

	case 'F':
		stack.PushFloat(slots.GetFloat(slotId))

	case 'J':
		stack.PushLong(slots.GetLong(slotId))

	case 'D':
		stack.PushDouble(slots.GetDouble(slotId))

	case 'L', '[':
		stack.PushRef(slots.GetRef(slotId))
	}
}
