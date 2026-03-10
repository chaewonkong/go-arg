package graph

func numIslands(grid [][]byte) int {
	yLen := len(grid)
	xLen := len(grid[0])

	cnt := 0
	for i := range yLen {
		for j := range xLen {
			if grid[i][j] == '1' {
				// call dfs
				dfs(j, i, grid)
				cnt++
			}
		}
	}

	return cnt
}

func dfs(x, y int, grid [][]byte) {
	yLen := len(grid)
	xLen := len(grid[0])

	if y < 0 || y >= yLen || x < 0 || x >= xLen || grid[y][x] == '0' {
		return
	}

	grid[y][x] = '0'
	directions := [][]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}

	for _, d := range directions {
		nx := x + d[0]
		ny := y + d[1]
		dfs(nx, ny, grid)
	}
}

func numIslandsBFS(grid [][]byte) int {
	rows := len(grid)
	cols := len(grid[0])
	directions := [][]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}

	cnt := 0
	for r := range rows {
		for c := range cols {
			if grid[r][c] == '1' {
				cnt++
				// add to queue
				queue := [][]int{{r, c}}
				grid[r][c] = '0'

				for len(queue) > 0 {
					cur := queue[0]
					queue = queue[1:]

					for _, d := range directions {
						nr := cur[0] + d[0]
						nc := cur[1] + d[1]

						if nr >= 0 && nc >= 0 && nr < rows && nc < cols && grid[nr][nc] == '1' {
							grid[nr][nc] = '0'
							queue = append(queue, []int{nr, nc})
						}
					}
				}
			}
		}
	}
	return cnt
}
