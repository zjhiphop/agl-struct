package sort

import "math"

func CountingSort(a []int, n int) {
	if n <= 1 {
		return
	}

	var times = 0

	var max int = math.MinInt32
	for i := range a {
		if a[i] > max {
			max = a[i]
		}
		times++
	}

	c := make([]int, max+1)
	for i := range a {
		c[a[i]]++
		times++
	}
	for i := 1; i <= max; i++ {
		c[i] += c[i-1]
		times++
	}

	r := make([]int, n)
	for i := range a {
		index := c[a[i]] - 1
		r[index] = a[i]
		c[a[i]]--
		times++
	}

	copy(a, r)

	println("Time O(", times, ")")
}

