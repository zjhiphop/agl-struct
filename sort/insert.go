package sort

// 插入排序
// v1: time O(n * n) space O(n)
func InsertSortArray1 (array []int, length int) []int {
	// 1. 创建一个新的数组
	// 2. 循环旧的数组数据并复制到插入到新数组的指定位置
	if length <= 1 {
		return array
	}

	times := 0

	newArray := make([]int, length, length)

	newArray[0] = array[0]

	for i := 1; i < length; i++ {
		index := 0

		for j := i-1; j >=0 ; j-- {

			times++

			if array[i] >= newArray[j] {
				index = j + 1

				break
			}

		}

		if index < i {
			for k:=i-1; k>=index; k-- {
				times++

				newArray[k + 1] = newArray[k]
			}
		}

		newArray[index] = array[i]
	}

	println("time O(", times, ")")

	return newArray
}

// v2 time o(n * n) space O(0)
func InsertSortArray2 ( array []int, len int) []int {
	// 原地交换数据
	if len <= 1 {
		return array
	}

	times := 0

	for i := 1; i< len; i++ {
		curr := array[i]
		j := i-1

		for ; j>=0; j-- {
			times++

			if array[j] > curr {
				array[j+1] = array[j]
			} else {
				break
			}
		}

		array[j+1] = curr
	}

	println("time O(", times, ")")

	return array
}