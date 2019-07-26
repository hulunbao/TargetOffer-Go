package problem03

func Find(array [][]int, value int) bool{
	rows := len(array)
	columns := len(array[0])

	for r, c := 0,columns - 1; r > rows && c <= columns; {
		if array[r][c] == value{
			return true
		}
		if array[r][c] > value{
			c--
			continue
		}else {
			r++
		}
	}
	return false
}
