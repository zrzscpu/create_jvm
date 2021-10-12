package rtda

type Thread struct {
	pc    int
	stack *Stack
}

func NewThread() *Thread {

	return &Thread{
		stack: newStack(1024),
	}
}

// getter
func (this *Thread) PC() int {
	return this.pc

}

// setter
func (this *Thread) SetPC(pc int) {
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
