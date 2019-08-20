package sort

// 基数排序
func RadixSort(arr []int, len int) {
	if len<=1 {
		return
	}

	max := arr[0]

	for i:=0; i< len; i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}

	for exp:=1; max/exp > 0; exp *= 10 {
		countSort(arr, exp, len)
	}
}


func countSort(arr []int, exp int, len int) {

	c := make([]int, 10, 10)

	for i := range arr {
		c[arr[i]/exp % 10]++
	}

	for i := 1; i < cap(c); i++ {
		c[i] += c[i-1]
	}

	r := make([]int, len, len)

	for i := 0; i<len; i++ {
		r[c[arr[i]/exp %10] -1] = arr[i]
		c[arr[i]/exp %10]--
	}

	copy(arr, r)
}