package sort

var exetimes int = 0

func MergeSort(arr []int, len int) {
	if len <=1 {
		return
	}

	sort(arr, 0, len-1)

	println("tims O(", exetimes, ")")
}

func sort(arr []int, start, end int) {
	if start >=end {
		return
	}

	mid := (start + end) /2

	sort(arr, start, mid)
	sort(arr, mid+1, end)
	merge(arr, start, mid, end)
}

func merge(arr []int, start, mid, end int) {
	temp := make([]int, end - start + 1)

	i := start
	j :=mid+1
	k := 0

	for ; i<=mid && j<= end; k++ {
		exetimes++

		if arr[i] > arr[j] {
			temp[k] = arr[j]
			j++
		} else {
			temp[k] = arr[i]
			i++
		}
	}

	for ;i<=mid;i++ {
		exetimes++

		temp[k] = arr[i]
		k++
	}

	for ;j<=end;j++ {
		exetimes++
		temp[k] = arr[j]
		k++
	}

	copy(arr[start:end+1], temp)
}
