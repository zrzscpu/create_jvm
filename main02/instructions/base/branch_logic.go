package base

import "create_jvm/main02/rtda"

//指令的跳转
func Branch(frame *rtda.Frame, offset int) {
	pc := frame.Thread().PC()

	nextPC := uint(int(pc) + offset)

	frame.SetNextPc(nextPC)

}
