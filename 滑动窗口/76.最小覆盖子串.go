package 滑动窗口

import "fmt"

func minWindow(s string, t string) string {
	m := map[byte]int{}
	for i := 0; i < len(t); i++ {
		m[t[i]]++
	}
	l := 0
	r := 0
	n := map[byte]int{}
	check := func() bool {
		for k, v := range m {
			if n[k] < v {
				return false
			}
		}
		return true
	}
	ansl := 0
	ansr := 0
	ans := 0
	for l <= r {

		for r < len(s) && !check() {
			n[s[r]]++
			r++
		}
		fmt.Println(n)
		for check() && l <= r {
			n[s[l]]--
			l++
		}
		if r-l+1 > ans {
			ans = r - l + 2
			ansl = l - 1
			ansr = r
		}
		if r == len(s)-1 {
			break
		}

	}
	return s[ansl : ansr+1]

}
