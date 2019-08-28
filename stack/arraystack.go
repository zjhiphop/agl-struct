package stack

type ArrayStack struct {
	data []string
	length int
	size int
}

func NewArrayStack(len int) *ArrayStack{
	if len <0 {
		return nil
	}

	return &ArrayStack{
		data: make([]string, len),
		length: len,
		size: 0,
	}
}

func (me *ArrayStack) push(data string) bool {
	if me.length == me.size {
		return false
	}

	me.data = append(me.data, data)

	me.size++

	return true
}

func (me *ArrayStack) pop() string{
	if me.size == 0 {
		return ""
	}

	item := me.data[me.size - 1]

	me.size--

	return item

}