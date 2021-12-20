package heap

//创建的对象,需要保留的信息:指向Class的指针
//该对象所拥有的属性
type Object struct {
	class *Class
	field Slots
}

func newObject(class *Class) *Object {
	return &Object{
		class: class,
		field: newSlots(class.instanceSlotCount),
	}
}

func (this *Object) Field() Slots {
	return this.field

}

func (this *Object) Class() *Class {
	return this.class

}

func (this *Object) IsInstanceOf(class *Class) bool {

	return class.isAssignableFrom(this.class)
}

//在java的构造方法中,会先对实例变量进行默认初始化,然后进行构造方法中的初始化过程
