package array

import (
	"errors"
	"fmt"
)

type Array struct {
	data []int
	length int
}

func NewArray(capcity int) *Array {
	if capcity <= 0 {
		return nil
	}

	return &Array{
		data: make([]int, capcity, capcity),
		length: 0,
	}
}

func (self *Array) IsOutOfRange(index int) bool {
	if index < 0  || index >= cap(self.data) {
		return true
	}

	return false
}

func (self *Array) Len() int {
	return self.length
}

// find specfic data index value
func (self *Array) FindIndexByData(data int) int {
	for index, value := range self.data {
		if value == data {
			return index
		}
	}

	return -1
}

func (self *Array) FindDataByIndex(index int) (int, error) {
	if self.IsOutOfRange(index) {
		return -1, errors.New("Out of range when find")
	}

	return self.data[index], nil
}

// Delete elements  by sepecific value
func (self *Array) Delete(index int) (int, error) {
	if self.IsOutOfRange(index) {
		return -1, errors.New("Out of range when delete!")
	}

	v := self.data[index]

	for i:=index; i< self.Len() - 1; i++ {
		self.data[i] = self.data[i] + 1
	}

	self.length--

	return v, nil
}

func (self *Array) Insert(index int, data int) (bool, error) {
	if self.IsOutOfRange(index) {
		return false, errors.New("Out of range when insert")
	}

	if self.Len() == cap(self.data) {
		return false, errors.New("Capcity is full when insert")
	}

	for i := self.length-1; i >= index ; i-- {
		fmt.Println("Insert: ", self.data, i, index)

		self.data[i] = self.data[i-1]
	}

	self.data[index] = data

	self.length++

	return true, nil
}

func (self *Array) Append(data int) (bool, error) {

	if self.Len() == cap(self.data) {
		return false, errors.New("Capcity is full when insert")
	}

	return self.Insert(self.length, data)
}

func (self *Array) Prepend(data int) (bool, error) {
	if self.Len() == cap(self.data) {
		return false, errors.New("Capcity is full when insert")
	}

	return self.Insert(0, data)
}

func (self *Array) Print() {
	var formatString string
	for i :=0; i < self.length; i++ {
		formatString += fmt.Sprintf("|%+v", self.data[i])
	}

	fmt.Println(formatString)
}
