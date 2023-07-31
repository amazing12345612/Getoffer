package main

func exist(board [][]byte, word string) bool {

	var dfs func(i, j int, z int) bool
	dfs = func(i, j int, z int) bool {
		if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || board[i][j] != word[z] || board[i][j] == '0' {
			return false
		}
		if z == len(word) {
			return true
		}
		board[i][j] = '0'
		res := dfs(i+1, j, z+1) || dfs(i-1, j, z+1) || dfs(i, j+1, z+1) || dfs(i, j-1, z+1)
		board[i][j] = word[z]
		return res

	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if word[0] == board[i][j] {
				if dfs(i, j, 0) {
					return true
				}
			}
		}
	}
	return false
}
