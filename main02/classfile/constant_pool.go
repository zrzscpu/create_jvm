package classfile

//用数组存储所有元素
type ConstantPool []ConstantInfo

//将常量池中的每一项进行解析
func readConstantPool(reader *ClassReader) ConstantPool {

	//常量池计数，并不是一个元素占据一个常量池索引，如
	//CONSTANT_Long_info和 CONSTANT_Double_info各占两个位置。
	cpCount := int(reader.readUint16())
	cp := make([]ConstantInfo, cpCount)
	//常量池0号位置表示不存方法任何常量引用
	for i := 1; i < cpCount; i++ {

		//读取当前元素的信息
		cp[i] = readConstantInfo(reader, cp)

		//这两种类型占据两个常量池引用
		//类型断言 x.(T) 其实就是判断 T 是否实现了 x 接口，
		//如果实现了，就把 x 接口类型具体化为 T 类型；

		//而 x.(type) 这种方式的类型断言，就只能和 switch 搭配使用，
		//因为它需要和多种类型比较判断，以确定其具体类型。
		//判断cp[i]的类型
		switch cp[i].(type) {
		//这两种类型占据两个位置的常量池索引
		case *ConstantLongInfo, *ConstantDoubleInfo:
			i++
		}
	}
	return cp
}

//按索引返回常量池中的元素
func (this ConstantPool) getConstantInfo(index uint16) ConstantInfo {

	if cpInfo := this[index]; cpInfo != nil {
		return cpInfo
	}
	//否则报错常量池中没有这个索引
	panic("Invalid constant pool index")
}

//为什么直接有了方法名和类型的信息,里面存的还是索引????

//按索引返回字段或方法的名字和描述符
func (this ConstantPool) getNameAndType(index uint16) (string, string) {

	//从常量池拿到名字和类型集合的信息
	ntInfo := this.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	//从里面提取
	name := this.getUtf8(ntInfo.nameIndex)
	_type := this.getUtf8(ntInfo.descriptorIndex)

	return name, _type

}

//从常量池中读取类名
func (this ConstantPool) getClassName(index uint16) string {

	//先找到描述类名的成员信息
	classInfo := this.getConstantInfo(index).(*ConstantClassInfo)
	return this.getUtf8(classInfo.nameIndex)
}

func (this ConstantPool) getUtf8(index uint16) string {

	utf8Info := this.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.str

}

//func (this *ConstantPool)readConstantInfo(reader *ClassReader, cp *ConstantPool) ConstantInfo {
//
//
//}
