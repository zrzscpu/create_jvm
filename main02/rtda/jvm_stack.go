package rtda

type Stack struct {
	maxSize uint
	size    uint
	//线程私有数据保存在栈帧中，栈帧通过一个链表连起来
	_top *Frame
}

func newStack(maxSize uint) *Stack {

	return &Stack{
		maxSize: maxSize,
	}
}

func (this *Stack) push(frame *Frame) {

	//当前栈的栈帧数超过上限
	if this.size >= this.maxSize {
		panic("java.lang.StackOverFlow")
	}
	if this._top != nil {
		frame.lower = this._top
	}
	this._top = frame
	this.size++
}

func (this *Stack) pop() *Frame {

	if this._top == nil {
		panic("jvm stack is empty!!!")
	}
	top := this._top
	this._top = top.lower
	//top弹出后不能还与更下面的栈帧有关联
	top.lower = nil
	this.size--
	return top
}

func (this *Stack) top() *Frame {

	if this._top == nil {
		panic("jvm stack is empty!!!")
	}
	return this._top
}
