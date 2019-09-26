package problem24

func isPostOrder(arr []int) bool {
	if arr == nil {
		return false
	}
	if len(arr) <= 2 {
		return true
	}
	root := arr[len(arr)-1]
	i := 0
	for ; i < len(arr)-1; i++ {
		if arr[i] > root {
			break
		}
	}
	j := i
	for ; j < len(arr)-1; j++ {
		if arr[j] < root {
			return false
		}
	}
	return isPostOrder(arr[:i]) && isPostOrder(arr[i:len(arr)-1])
}
