package main

func hammingWeight(num uint32) int {
	count := 0
	if num == 1 {
		return 1
	}
	for num > 1 {
		if num&1 == 1 {
			count++
		}
		num = num >> 1
	}
	return count
}
