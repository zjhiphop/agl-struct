package sort

/*
排序过程：

j=0, 1, i=0 1
1      111 23 31 43  54 45
> i=1

j=1 111, i=1 1111
1      111    23 31 43  54 45

j=2 23, i=1 111
1      23    111 31 43  54 45
> i=2 111

j = 3   31
1      23   31 111 43  54 45
> i = 3 111

j = 4, 43
1      23   31 43 111  54 45
> i = 4 111

j = 5, 54
1      23   31 43 111  54 45
> i=4 111

1      23   31 43 45  54 111
> i =4, j =5

*/
func QuickSort(arr []int, len int) {
	separate(arr, 0, len-1)

	println("time O(", times, ")")
}

func separate(arr []int, start,end int) {
	if start >= end {
		return
	}

	i := partition(start, end, arr)

	separate(arr, start, i-1)
	separate(arr, i+1, end)
}

var times int = 0

func partition(start int, end int, arr []int) int {
	pivot := arr[end]

	i := start

	for j :=start; j< end; j++ {
		times++

		if arr[j] < pivot {
			if i != j {
				arr[i], arr[j] = arr[j], arr[i]
			}

			i++
		}
	}

	arr[i], arr[end] = arr[end], arr[i]

	return i
}