package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	k := (m + n) / 2
	l := 0
	r := 0
	a, b := -1, -1
	for i := 0; i <= k; i++ {
		a = b
		if l < m && (r >= n || nums1[l] < nums2[r]) {
			b = nums1[l]
			l++
		} else {
			b = nums2[r]
			r++
		}
	}
	if (m+n)%2 == 1 {
		return float64(b)
	}
	return (float64(a) + float64(b)) / 2
}
