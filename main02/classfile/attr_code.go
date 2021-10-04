package classfile

//Code_attribute {
//	u2 attribute_name_index;
//	u4 attribute_length;
//	u2 max_stack; u2 max_locals;
//	u4 code_length;
//	u1 code[code_length];
//	u2 exception_table_length;
//	{
//		u2 start_pc;
//		u2 end_pc;
//		u2 handler_pc;
//		u2 catch_type;
//	}
//	exception_table[exception_table_length];
//	u2 attributes_count;
//	attribute_info
//		attributes[attributes_count];
//	}

type CodeAttribute struct {
	cp             ConstantPool
	maxStack       uint16
	maxLocals      uint16
	code           []byte
	exceptionTable []*ExceptionTableEntry
	attributes     []AttributeInfo
}

type ExceptionTableEntry struct {
	startPc   uint16
	endPc     uint16
	handlerPc uint16
	catchType uint16
}

func (this *CodeAttribute) readInfo(reader *ClassReader) {
	//操作数栈最大容量
	this.maxStack = reader.readUint16()
	//本地变量表容量
	this.maxLocals = reader.readUint16()

	//返回字节码指令的长度
	codeLength := reader.readUint32()
	//读取字节码指令
	this.code = reader.readBytes(codeLength)

	//解析异常表
	this.exceptionTable = readExceptionTable(reader)

	//解析属性表
	this.attributes = readAttributes(reader, this.cp)

}

//解析属性表的函数
func readExceptionTable(reader *ClassReader) []*ExceptionTableEntry {
	//异常表的长度
	exceptionTableLength := reader.readUint16()

	//异常表分配空间
	exceptionTable := make([]*ExceptionTableEntry, exceptionTableLength)
	//
	for i := range exceptionTable {
		exceptionTable[i] = &ExceptionTableEntry{
			startPc: reader.readUint16(),
			endPc:   reader.readUint16(),

			handlerPc: reader.readUint16(),
			catchType: reader.readUint16(),
		}
	}
	return exceptionTable
}
