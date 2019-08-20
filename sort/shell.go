package sort

func ShellSort(arr []int, len int) {
	if len <=1 {
		return
	}

	times := 0

	gap := len / 2

	for gap >=1 {
		j := 0

		for i:=gap; i<len; i++ {
			curr := arr[i]

		//	对第i个元素以及之前相同的gap间距元素进行对比排序
		   for j = i - gap; j>=0 && curr < arr[j]; j -= gap {

			   times++
			   arr[ j + gap ] = arr[j]
		   }

		   arr[j + gap] = curr
		}

		gap /= 2
	}

	println("time O(", times, ")")
}
