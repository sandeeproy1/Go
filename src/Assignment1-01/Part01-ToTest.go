func CountIslands(grid [][]int) int {
	var count int
	count = 0
	rowNo := len(grid)
	colNo := len(grid[0])
	for i := 0; i < rowNo; i++ {
		for j := 0; j < colNo; j++ {
			if grid[i][j] == 1 {
				count = count + 1
				checkIsland(i, j, grid)
			}
		}
	}
	return count
}

func checkIsland(a int, b int, array [][]int) {
	if a < 0 || a >= len(array) || b < 0 || b >= len(array[a]) || array[a][b] != 1 {
		return
	}
	array[a][b] = -1
	checkIsland(a, b-1, array)
	checkIsland(a, b+1, array)
	checkIsland(a-1, b, array)
	checkIsland(a+1, b, array)
}