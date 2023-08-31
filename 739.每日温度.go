package main

func dailyTemperatures(temperatures []int) []int {
	s := []int{len(temperatures) - 1}
	ans := []int{0}
	for i := len(temperatures) - 2; i >= 0; i-- {
		top := s[len(s)-1]
		for len(s) > 0 && temperatures[i] > temperatures[top] {
			s = s[:len(s)-1]
		}
		s = append(s, i)
		if len(s) == 0 {
			ans[i] = 0
		} else {
			ans[i] = s[len(s)-1] - i
		}

	}
	return ans
}
