package search

func BSearch(arr []int, len int, val int) int {
  low := 0
  high := len -1

  for low <= high {
  	mid := low + (high-low) >> 1

  	if arr[mid] == val {
  		return mid
	} else if arr[mid] > val {
		high = mid -1
	} else {
		low = mid + 1
	}
  }

  return -1
}

func BSearch2(arr []int, len int, val int) int {
	return partition(arr, 0, len-1, val)
}

func partition(arr []int, low int, high int, val int) int{
	mid := low + (high - low)>>2

	if arr[mid] == val {
		return mid
	} else if arr[mid] > val {
		return partition(arr, low, mid -1, val)
	} else {
		return partition(arr, mid + 1, high, val)
	}
}