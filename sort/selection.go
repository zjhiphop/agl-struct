package sort

func SelectionSort(arr []int, len int) {
	if len <=1 {
		return
	}

	times := 0

	for i:=0; i< len; i++ {
		minIndex := i

		for j:= i+1; j < len; j++ {

			times++

			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}

		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}

	println(times)
}
