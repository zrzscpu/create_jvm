package rtda

import "create_jvm/main02/rtda/heap"

type Thread struct {
	pc    uint
	stack *Stack
}

func NewThread() *Thread {

	return &Thread{
		stack: newStack(1024),
	}
}

// getter
func (this *Thread) PC() uint {
	return this.pc

}

// setter
func (this *Thread) SetPC(pc uint) {
	this.pc = pc
}

func (this *Thread) PushFrame(frame *Frame) {

	this.stack.push(frame)

}

func (this *Thread) PopFrame() *Frame {

	return this.stack.pop()
}

func (this *Thread) CurrentFrame() *Frame {

	return this.stack.top()
}

//func (this *Thread) NewFrame(maxLocals, maxStack uint16) *Frame {
//	return NewFrame(this, maxLocals, maxStack)
//}

func (this *Thread) NewFrame(method *heap.Method) *Frame {
	return newFrame(this, method)
}

func (this *Thread) SetNextPc(nextPC int) {

}
