package sort

func getMax(arr []int) int {
	max := 0

	for i := range arr {
		if i > max {
			max = i
		}
	}

	return max
}

func BucketSort(arr []int, arrLen int) {
	if arrLen <= 1 {
		return
	}

	max := getMax(arr)
	buckets := make([][]int, arrLen / 10) // seperate to length/10 buckets

	for i := 0; i < arrLen; i++ {
		index := arr[i] * ( arrLen - 1 ) / (10 * max) // sperate data to different bucket

		buckets[index] = append(buckets[index], arr[i])
	}

	datalen := 0 // mark current data processed

	for j := 0; j < len(buckets); j++ {
		bucLen := len(buckets[j])

		if bucLen > 0 {
			// use quicksort to sort every bucket
			QuickSort(buckets[j], bucLen)

			copy(arr[datalen:], buckets[j])

			datalen += bucLen
		}
	}
}
