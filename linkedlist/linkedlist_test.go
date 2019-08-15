package linkedlist

import (
	"testing"
)

func TestLinkedList_InsertAfter(t *testing.T) {
	list := NewLinkedList()

	err := list.InsertAfter(list.head, 123)

	if err != nil {
		t.Fatal(err)
	}

	list.Print()
}

func TestLinkedList_FindByIndex(t *testing.T) {
	list := NewLinkedList()

	for i:=0; i< 10 ; i++ {
		err := list.Append(i)

		if err != nil {
			t.Fatal(err)
		}
	}

	t.Log(list.FindByIndex(3))
	t.Log(list.FindByIndex(5))
	t.Log(list.FindByIndex(8))

}