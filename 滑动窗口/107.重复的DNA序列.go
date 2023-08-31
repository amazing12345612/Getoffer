package 滑动窗口

func findRepeatedDnaSequences(s string) (ans []string) {
	if len(s) < 10 {
		return
	}
	l := 0
	r := 9
	m := map[string]int{}
	for r < len(s) {
		ts := s[l : r+1]
		m[ts]++
		if m[ts] == 1 {
			ans = append(ans, ts)
		}
		l++
		r++
	}
	return ans

}
