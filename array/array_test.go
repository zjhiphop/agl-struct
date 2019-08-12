package array

import (
	"fmt"
	"testing"
)

func TestInsert(t *testing.T) {
	arr := NewArray(10)

	fmt.Println(cap(arr.data), arr.length)

	for i:=0; i< 8; i++ {

		fmt.Println("Before insert: ", i)

		isSuccess, err := arr.Insert(i, i)

		if !isSuccess {
			t.Fatal(err)
		}

	}

	arr.Print()

	result, err := arr.Append(4)

	if !result {
		t.Error(err)
	}

	arr.Print()

	result, err = arr.Append(1323)


	if !result {
		t.Error(err)
	}
}