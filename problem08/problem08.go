package problem08

func minNumberInrotateArray(arr []int) int {
	if arr == nil {
		return -2
	}
	start := 0
	end := len(arr) - 1
	mid := 0
	for arr[start] >= arr[end] {
		if end-start == 1 {
			return arr[end]
		}
		mid = start + (end-start)/2
		if arr[start] == arr[end] && arr[mid] == arr[start] {
			return MinInOrder(arr, start, end)
		}

		if arr[mid] >= arr[start] {
			start = mid
		} else {
			end = mid
		}
	}
	return arr[mid]
}

// MinInOrder 顺序查找
func MinInOrder(arr []int, start, end int) int {
	result := start
	value := arr[start]
	for i := start + 1; i <= end; i++ {
		if arr[i] < value {
			result = i
			value = arr[i]
		}
	}
	return arr[result]
}
