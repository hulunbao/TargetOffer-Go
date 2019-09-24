package problem20

// PrintMatrixClockwisely 顺时针打印矩阵
func PrintMatrixClockwisely(matrix [][]int) []int {
	var res []int
	row := len(matrix)
	if row == 0 {
		return res
	}
	column := len(matrix[0])

	// 计算元素个数
	count := row * column
	// 第几层
	layer := 0

	// 初始边界
	startRow, endRwo, startCol, endCol := 0, 0, 0, 0
	for count > 0 {
		startRow, endRwo = layer, row-layer-1
		startCol, endCol = layer, column-layer-1
		for i := startCol; i <= endCol && count > 0; i++ {
			res = append(res, matrix[startRow][i])
			count--
		}
		for i := startRow + 1; i <= endRwo && count > 0; i++ {
			res = append(res, matrix[i][endCol])
			count--
		}
		for i := endCol - 1; i >= startCol && count > 0; i-- {
			res = append(res, matrix[endRwo][i])
			count--
		}
		for i := endRwo - 1; i > startRow && count > 0; i-- {
			res = append(res, matrix[i][startCol])
			count--
		}
		layer++
	}

	return res

}
