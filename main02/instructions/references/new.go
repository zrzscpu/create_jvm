package references

import (
	"create_jvm/main02/instructions/base"
	"create_jvm/main02/rtda"
	"create_jvm/main02/rtda/heap"
)

//从运行时常量池中取得一个类的符号引用,然后解析成Class对象根据class对象创建 实例对象
type NEW struct {
	base.Index16Instruction
}

func (this *NEW) Execute(frame *rtda.Frame) {
	cp := frame.Method().Class().ConstantPool()
	classRef := cp.GetConstant(this.Index).(*heap.ClassRef)

	//通过引用解析得到Class对象
	class := classRef.ResolvedClass()
	//不是接口，不是抽象类才可以创建对象
	if !class.IsInterface() && !class.IsAbstract() {
		object := class.NewObject()
		frame.OperandStack().PushRef(object)
	} else {
		panic("java.lang.InstantiationError")
	}
}
