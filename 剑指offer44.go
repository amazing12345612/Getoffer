package main

func findNthDigit(n int) int {
	i := 1
	if n < 10 {
		return n
	}
	n -= 10
	i += 1
	t := 9
	for n-t*10*i > 0 {
		n -= t * 10 * i
		t = 9 * 10
		i += 1
	}
	a := n % i
	if a == 0 {
		return n / i % 9
	} else if a == 1 {
		return n
	}
	return 0
}
