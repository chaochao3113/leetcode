/**
* @Author: Chao
* @Date: 2020-04-20 18:50
* @Version: 1.0
 */

package main

import "fmt"

func main() {
	grid := [][]byte{{'1','1','0','0','0'}, {'1','1','0','0','0'}, {'0','0','1','0','0'}, {'0','0','0','1','1'}}
	fmt.Println(numIslands(grid))
}

func numIslands(grid [][]byte) int {
	a := len(grid)
	if a == 0 {
		return 0
	}

	b := len(grid[0])

	num_islands := 0
	for i := 0; i < a; i++ {
		for j := 0; j < b; j++ {
			if grid[i][j] == '1' {
				num_islands++
				dfs(grid, i, j)
			}
		}
	}
	return num_islands
}
//深度优先搜索
func dfs(grid [][]byte, a, b int) {
	aa := len(grid)
	bb := len(grid[0])

	grid[a][b] = '0'
	if b+1 < bb && grid[a][b+1] == '1' {
		grid[a][b+1] = '0'
		dfs(grid, a, b+1)
	}
	if b-1 >= 0 && grid[a][b-1] == '1' {
		grid[a][b-1] = '0'
		dfs(grid, a, b-1)
	}
	if a+1 < aa && grid[a+1][b] == '1' {
		grid[a+1][b] = '0'
		dfs(grid, a+1, b)
	}
	if a-1 >= 0 && grid[a-1][b] == '1' {
		grid[a-1][b] = '0'
		dfs(grid, a-1, b)
	}
}