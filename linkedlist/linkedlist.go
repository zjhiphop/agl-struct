package linkedlist

import (
	"fmt"
	"github.com/pkg/errors"
)

type Node struct {
	next *Node
	data interface{}
}

type LinkedList struct {
	head *Node
	length uint
}

func NewNode(value interface{}) *Node {
	return &Node{nil, value}
}

func NewLinkedList() *LinkedList{
	return &LinkedList{NewNode(0), 0}
}

func (self *Node) GetNext(node *Node) *Node {
	if nil == node {
		return nil
	}

	return node.next
}

func (self *Node) GetValue(node *Node) interface{} {
	if nil == node {
		return nil
	}

	return node.data
}

// insert new node after specific node
func (self *LinkedList) InsertBefore(node *Node, v interface{}) error {
	if nil == node {
		return errors.New("Try to insert before a empty node")
	}

	newNode := NewNode(v)
	curr := self.head.next
	pre := self.head

	for nil != curr {
		if curr == node {
			break
		}

		pre = curr
		curr = curr.next
	}

	if nil == curr {
		return errors.New("Can not find node to insert")
	}

	pre.next = newNode
	newNode.next = curr

	self.length++

	return nil
}

func (self *LinkedList) InsertAfter(node *Node, v interface{}) error {
	if nil == node {
		return errors.New("Try to insert after a empty node")
	}

	curr := self.head

	for nil != curr {
		if curr == node {
			break
		}

		curr = curr.next
	}

	if curr == nil {
		return errors.New("Can not find node to insert")
	}

	newNode := NewNode(v)
	nextNode := curr.next
	curr.next = newNode
	newNode.next = nextNode

	self.length++

	return nil
}

func (self *LinkedList) Prepend(v interface{}) error {
	return self.InsertAfter(self.head, v)
}

func (self *LinkedList) Append(v interface{}) error {
	newNode := NewNode(v)

	head := self.head

	if nil == head.next {

		head.next = newNode

	} else {

		curr := head

		for nil != curr.next {
			curr = curr.next
		}

		curr.next = newNode
	}

	self.length++

	return nil
}


// find
func (self *LinkedList) FindByIndex(index uint) (*Node, error) {

	if index > self.length {
		return nil, errors.New("Find out of range")
	}

	curr := self.head

	for i := uint(0); i< index; i++ {
		curr = curr.next
	}

	return curr, nil
}

func (self *LinkedList) FindByData(data interface{}) (*Node, error) {
	curr := self.head

	for nil != curr {
		curr = curr.next
	}

	return curr, nil
}

func (self *LinkedList) DelNode(node *Node) (bool, error) {
	if nil == node {
		return false, errors.New("Try to delete empty node")
	}

	curr := self.head

	for nil != curr {

		if(curr.next == node) {
			curr.next = curr.next.next

			return true, nil
		}

		curr = curr.next
	}

	return false, nil
}

func (self *LinkedList) Print() {
	curr := self.head

	var formatStr string

	for nil != curr.next {
		curr = curr.next

		formatStr += fmt.Sprintf("|%+v", curr.data)
	}

	fmt.Println(formatStr)
}
