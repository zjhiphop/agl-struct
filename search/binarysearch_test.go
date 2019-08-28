package search

import (
	"../array"
	"testing"
)

func TestBSearch(t *testing.T) {
	length := 20

	arr := array.NewArray(length)

	for i := 0; i < length; i++ {
		arr.Append(i)
	}

	println(BSearch(arr.Data(), arr.Len(), 40))

}