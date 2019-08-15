package sort

func BubleSort(arr []int, len int) {
	if len <=1 {
		return
	}

	flag := false

	times := 0

	for i :=0; i< len; i++ {
		for j :=0; j< len-i-1; j++ {

			times++

			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]

				flag = true
			}
		}

		if !flag {
			break
		}
	}

	println("time:", times)
}