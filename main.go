package main

import (
	"./array"
	"./sort"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	length := 100
	arr := array.NewArray(length)
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	for i:=0; i< length; i++ {
		arr.Append(r1.Intn(length))
	}

	sort.MergeSort(arr.Data(), arr.Len())

	fmt.Println(arr.Data())
}
