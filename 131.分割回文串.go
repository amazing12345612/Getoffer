package main

func partition(s string) [][]string {

	ans := make([][]string, 0)
	path := make([]string, 0)

	var dfs func(i int)
	dfs = func(i int) {
		if i == len(s) {
			ans = append(ans, append([]string{}, path...))
			return
		}

		for j := i + 1; j <= len(s); j++ {
			if !check(s[i:j]) {
				continue
			}
			path = append(path, s[i:j])
			dfs(j + 1)
			path = path[:len(path)-1]

		}

	}
	dfs(0)
	return ans

}
func check(s string) bool {
	l := 0
	r := len(s) - 1
	for l <= r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--

	}
	return true

}
