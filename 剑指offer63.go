package main

func maxProfit(prices []int) int {
	Min := prices[0]
	res := 0
	for _, v := range prices {
		if v < Min {
			Min = v
		} else {
			if v-Min > res {
				res = v - Min
			}
		}
	}
	return res
}
