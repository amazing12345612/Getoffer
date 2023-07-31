package main

func fib(n int) int {
	if n == 0 {
		return 0
	}
	a := 0
	b := 1
	for i := 1; i < n; i++ {
		t := b
		b = (a + b) % (1e10 + 7)
		a = t % (1e10 + 7)
	}
	return b

}
