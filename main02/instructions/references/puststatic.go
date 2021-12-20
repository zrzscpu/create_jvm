package references

import (
	"create_jvm/main02/instructions/base"
	"create_jvm/main02/rtda"
	"create_jvm/main02/rtda/heap"
)

//给类的某个静态变量赋值，从操作数栈中取待放入的值
type PUT_STATIC struct {
	base.Index16Instruction
}

func (this *PUT_STATIC) Execute(frame *rtda.Frame) {
	//当前执行中的method，class，class的运行时常量池
	curmethod := frame.Method()
	curClass := curmethod.Class()
	cp := curClass.ConstantPool()
	//获取当前的域对象
	ref := cp.GetConstant(this.Index).(*heap.FieldRef)
	field := ref.ResolvedField()

	//域对象所属类的Class对象
	class := field.Class()

	if !field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	//如果是final static字段，则实际操作的 是静态常量，只能在clinit给它赋值
	if field.IsFinal() {
		if curClass != class || curmethod.Name() != "<clinit>" {
			panic("java.lang.IllegalAccessError")
		}
	}

	descriptor := field.Descriptor()
	id := field.SlotId()
	//静态字段
	slots := class.StaticVars()
	stack := frame.OperandStack()

	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I':
		slots.SetInt(id, stack.PopInt())

	case 'F':
		slots.SetFloat(id, stack.PopFloat())

	case 'J':
		slots.SetLong(id, stack.PopLong())

	case 'D':
		slots.SetDouble(id, stack.PopDouble())

	case 'L':
		slots.SetRef(id, stack.PopRef())
	}

}
