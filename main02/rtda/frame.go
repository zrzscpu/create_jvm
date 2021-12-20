package rtda

import "create_jvm/main02/rtda/heap"

type Frame struct {
	lower        *Frame
	localVars    LocalVars
	operandStack *OperandStack
	thread       *Thread
	nextPc       uint
	method       *heap.Method
}

//func NewFrame(thread *Thread, maxLocals, maxStack uint16) *Frame {
//
//	return &Frame{
//		thread:       thread,
//		localVars:    newLocalVars(maxLocals),
//		operandStack: newOperandStack(maxStack),
//	}
//}
func newFrame(thread *Thread, method *heap.Method) *Frame {
	return &Frame{
		thread:       thread,
		localVars:    newLocalVars(method.MaxLocals()),
		operandStack: newOperandStack(method.MaxStack()),
		method:       method,
	}
}

func (this *Frame) LocalVars() LocalVars {

	return this.localVars
}

func (this *Frame) OperandStack() *OperandStack {

	return this.operandStack
}

func (this *Frame) Thread() *Thread {
	return this.thread
}

func (this *Frame) NextPc() uint {
	return this.nextPc
}

func (this *Frame) SetNextPc(nextPc uint) {
	this.nextPc = nextPc
}

func (this *Frame) Method() *heap.Method {
	return this.method
}
