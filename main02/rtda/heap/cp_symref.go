package heap

//符号引用的父类
type SymRef struct {
	cp        *ConstantPool
	className string
	//如果当前的引用不是Classref ， 那就保存所属Class对象（解析后的）
	class *Class
}

//将符号引用解析为Class对象
func (this *SymRef) ResolvedClass() *Class {
	if this.class == nil {
		this.resolveClassRef()
	}
	return this.class
}

func (this *SymRef) resolveClassRef() {
	//引用当前这个类的 类对象
	class := this.cp.class

	//使用当前引用这个属性的类的类加载器进行记载
	loadClass := class.loader.LoadClass(this.className)

	//如果引用这个属性的类 无权访问这个 属性所属的类
	if !loadClass.isAccessibleTo(class) {
		panic("java.lang.IllegalAccessError")
	}
	this.class = loadClass
}
