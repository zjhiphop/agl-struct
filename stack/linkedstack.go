package stack

import "fmt"

type node struct {
	 next *node
	 val interface{}
}

type LinkedListStack struct {
	top *node
}

func NewLinkedListStack() *LinkedListStack {
	return &LinkedListStack{nil}
}

func (me LinkedListStack) IsEmpty() bool {
	return me.top == nil
}

func (me *LinkedListStack) Push(data interface{}) bool {
	nd := &node{val: data, next: nil}

	if me.IsEmpty() {
		me.top = nd
	} else {
		nd.next = me.top
		me.top = nd.next
	}

	return true
}

func (me *LinkedListStack) Pop() interface{} {
	if me.IsEmpty() {
		return nil
	}

	result := me.top.val
	me.top = me.top.next

	return result
}

func (me *LinkedListStack) Top() interface{} {
	if me.IsEmpty() {
		return nil
	}

	return me.top.val
}

func (me *LinkedListStack) Print() {
	if me.IsEmpty() {
		return
	} else {
		curr := me.top

		for nil != curr {
			fmt.Println(curr.val)
			curr = curr.next
		}
	}

}
