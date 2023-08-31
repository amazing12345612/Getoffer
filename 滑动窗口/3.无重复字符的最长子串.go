package 滑动窗口

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	m := map[byte]bool{}
	r := 0
	l := -1
	ans := 0
	for ; l < len(s)-1; l++ {
		if l != -1 {
			delete(m, s[l])
		}

		for r < len(s)-1 && !m[s[r]] {
			m[s[r]] = true
			r++
		}
		if r-l+1 > ans {
			ans = r - l + 1
		}

	}
	return ans

}
