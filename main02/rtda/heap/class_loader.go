package heap

import (
	"create_jvm/main02/classfile"
	"create_jvm/main02/classpath"
	"fmt"
)

//根据classpath 来读取class文件,加载后的class文件指针保存在classmap
type ClassLoader struct {
	cp       *classpath.ClassPath
	classMap map[string]*Class
}

func NewClassLoader(cp *classpath.ClassPath) *ClassLoader {
	return &ClassLoader{
		cp:       cp,
		classMap: make(map[string]*Class),
	}
}

//将name指定的类加载到方法区
func (this *ClassLoader) LoadClass(name string) *Class {
	//如果存在ok里面就是true
	if class, ok := this.classMap[name]; ok {
		return class
	}
	return this.loadNonArrayClass(name)
}

//从内存中读取并链接
func (self *ClassLoader) loadNonArrayClass(name string) *Class {
	//读取字节码
	data, entry := self.readClass(name)

	//转为classfile对象,加载父类和接口
	class := self.defineClass(data)

	link(class)
	fmt.Printf("[Loaded %s from %s]\n", name, entry)
	return class

}

//读取类，返回类的字节码和加载地址
func (this *ClassLoader) readClass(name string) ([]byte, classpath.Entry) {
	data, entry, err := this.cp.ReadClass(name)
	if err != nil {
		panic(("java.lang.ClassNotFoundException: " + name))
	}
	return data, entry

}

//将字节码解析成类文件对象，再解析成Class对象
func (this *ClassLoader) parseClass(data []byte) *Class {
	cf, err := classfile.Parse(data)
	if err != nil {
		panic("java.lang.ClassFormatError")
	}
	return newClass(cf)
}

//其实就是给当前类设置成是当前类加载器加载的,并加载父类,和接口
func (this *ClassLoader) defineClass(data []byte) *Class {
	class := this.parseClass(data)
	//设置成是当前类加载器进行加载的
	class.loader = this
	//加载父类
	resolveSuperClass(class)
	//加载接口
	resolveInterfaces(class)
	this.classMap[class.name] = class
	return class
}

//加载父类对象
func resolveSuperClass(class *Class) {
	if class.name != "java/lang/Object" {
		class.loader.LoadClass(class.superClassName)
	}
}

//为加载后创建的Class对象解析所有的接口
func resolveInterfaces(class *Class) {
	length := len(class.interfaceNames)
	if length > 0 {
		class.interfaces = make([]*Class, length)
		for i, interfaceName := range class.interfaceNames {
			class.interfaces[i] = class.loader.LoadClass(interfaceName)
		}
	}
}

func link(class *Class) {
	verify(class)

	prepare(class)
}
func verify(class *Class) {
	//先不做验证
}

//计算类变量和实例变量的个数，并分配内存
func prepare(class *Class) {
	calcInstanceFieldSlotIds(class)
	calcStaticFieldSlotIds(class)

	//静态初始化
	allocAndInitStaticVars(class)
}

//计算实例字段的个数并编号
func calcInstanceFieldSlotIds(class *Class) {
	slotid := uint(0)
	//先计算父类的实例字段的个数,这里父类的Class对象还未解析成功???? 统计情况不会出错?
	if class.superClass != nil {
		slotid = class.superClass.instanceSlotCount
	}
	for _, field := range class.fields {
		if !field.IsStatic() {
			//给field添加slot
			field.slotId = slotid
			slotid++
			//如果是LongOrDouble，长度加1
			if field.IsLongOrDouble() {
				slotid++
			}
		}
	}
	class.instanceSlotCount = slotid
}

//计算静态字段的个数并编号
func calcStaticFieldSlotIds(class *Class) {
	slotid := uint(0)
	//在子类中会继承父类的静态字段//本质上是如何访问的;---------------------------------------------------------------
	for _, field := range class.fields {
		if field.IsStatic() {
			//给field添加slot
			field.slotId = slotid
			slotid++
			//如果是LongOrDouble，长度加1
			if field.IsLongOrDouble() {
				slotid++
			}
		}
	}
	class.staticSlotCount = slotid
}

//统计静态变量和实例变量的个数  并分配空间赋初值
func allocAndInitStaticVars(class *Class) {
	class.staticVars = newSlots(class.staticSlotCount)
	for _, field := range class.fields {
		//如果是常量的在这里要完成一次初始化
		if field.IsStatic() && field.IsFinal() {
			initStaticFinalVar(class, field)
		}
	}
}

// 为static final修饰的常量的值在编译期就可知道,在这里直接赋值
func initStaticFinalVar(class *Class, field *Field) {
	vars := class.staticVars
	cp := class.constantPool
	cpIndex := field.ConstantValueIndex()
	slotId := field.SlotId()
	if cpIndex > 0 {
		switch field.descriptor {

		//z	bool
		case "Z", "B", "C", "S", "I":
			val := cp.GetConstant(cpIndex).(int32)
			vars.SetInt(slotId, val)

		//long
		case "J":
			val := cp.GetConstant(cpIndex).(int64)
			vars.SetLong(slotId, val)

		case "F":
			val := cp.GetConstant(cpIndex).(float32)
			vars.SetFloat(slotId, val)

		case "D":
			val := cp.GetConstant(cpIndex).(float64)
			vars.SetDouble(slotId, val)

		case "Ljava/lang/String;":
			panic("todo") // 在第

		}
	}
}
