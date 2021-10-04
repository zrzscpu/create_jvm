package classfile

/*属性表的通用定义
attribute_info {
	u2 attribute_name_index;
	u4 attribute_length;
	u1 info[attribute_length];
}
*/
type AttributeInfo interface {
	readInfo(reader *ClassReader)
}

//读取所有的属性表
func readAttributes(reader *ClassReader, cp ConstantPool) []AttributeInfo {
	//读取长度
	attributesCount := reader.readUint16()
	attributes := make([]AttributeInfo, attributesCount)

	for i, _ := range attributes {
		attributes[i] = readAttribute(reader, cp)
	}
	return attributes
}

//读取单个attribute
func readAttribute(reader *ClassReader, cp ConstantPool) AttributeInfo {

	//属性的名称
	attrNameIndex := reader.readUint16()
	attrName := cp.getUtf8(attrNameIndex)

	//属性的长度(属性也是一张表)
	attrLen := reader.readUint32()
	//创建这个表
	attrInfo := newAttributeInfo(attrName, attrLen, cp)

	//读取这个表的信息
	attrInfo.readInfo(reader)
	return attrInfo
}

func newAttributeInfo(attrName string, attrLen uint32, cp ConstantPool) AttributeInfo {
	switch attrName {
	//code属性表
	case "Code":
		return &CodeAttribute{cp: cp}

	//常量表
	case "ConstantValue":
		return &ConstantValueAttribute{}

	case "Deprecated":
		return &DeprecatedAttribute{}

	//异常表
	case "Exceptions":
		return &ExceptionsAttribute{}

	//本地变量表
	//case "LocalVariableTable": return &LocalVariableTableAttribute{}

	//class文件属性表
	case "SourceFile":
		return &SourceFileAttribute{cp: cp}

	default:
		return &UnparsedAttribute{attrName, attrLen, nil}
	}
}
