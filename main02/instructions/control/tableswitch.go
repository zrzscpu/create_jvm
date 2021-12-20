package control

import (
	"create_jvm/main02/instructions/base"
	"create_jvm/main02/rtda"
)

/**先这样的可以变成Switchtable
int chooseNear(int i) {
	switch (i) {
		case 0:
			return 0;
		case 1:
			return 1;
		case 2:
			return 2;
		default:
			return -1;
				}
}
*/

//很显然这种Switchtable的方式比Switchcase的方式快的多

//switch case 指令跳转,通过index 在jumpOffsets匹配跳转的路径
//每个分支跳转的字节码偏移量从字节码中读取,分支的个数是high-low+1
type TABLE_SWITCH struct {
	defaultOffset int32
	//high-low 指定这个偏移表的大小
	low         int32
	high        int32
	jumpOffsets []int32
}

//tableswitch指令操作码的后面有0~3字节的padding
//以保证 defaultOffset在字节码中的地址是4的倍数
func (this *TABLE_SWITCH) FetchOperands(reader *base.BytecodeReader) {
	reader.SkipPadding()
	this.defaultOffset = reader.ReadInt32()
	this.low = reader.ReadInt32()
	this.high = reader.ReadInt32()
	jumpOffsetsCount := this.high - this.low + 1
	this.jumpOffsets = reader.ReadInt32s(jumpOffsetsCount)
}

func (this *TABLE_SWITCH) Execute(frame *rtda.Frame) {
	//从操作数栈中弹出offset 的 索引
	index := frame.OperandStack().PopInt()
	var offset int
	//先判断索引是否在这个偏移表的索引范围内,不在的话就是使用默认的跳转
	if index >= this.low && index <= this.high {
		offset = int(this.jumpOffsets[index-this.low])
	} else {
		offset = int(this.defaultOffset)
	}
	base.Branch(frame, offset)
}
