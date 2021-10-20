package rtda

type Slot struct {

	//存放一个值或一个引用
	/**
	选择int32的原因，两个slot存放一个long或者double，一个单位的slot是32位
	为什么不用uint32？
	使用uint32 ，对于存放一个int64的值，需要转换两次，而使用int32无需转换
	*/
	num int32
	ref *Object
}
