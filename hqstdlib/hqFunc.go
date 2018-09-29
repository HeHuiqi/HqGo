package hqstdlib

type HqCount int

//值传递
func (c HqCount) HqValuePassAdd(add HqCount){
	c += add
}
//指针传递（修改值时使用指针传递）
func (c *HqCount)HqPointerPassAdd(add HqCount)  {
	*c += add
}
