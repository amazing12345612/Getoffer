package main

func longestValidParentheses(s string) int {
	p := make([]int, 0)
	ans := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			p = append(p, i)
		} else {
			p = p[:len(p)-1]
			if len(p) == 0 {
				p = append(p, i)
			} else {
				if i-p[len(p)-1] > ans {
					ans = i - p[len(p)-1]
				}
			}
		}
	}

	return ans
}
