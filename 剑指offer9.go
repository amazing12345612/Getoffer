package main

type CQueue struct {
	a []int
	b []int
}

func Constructor() CQueue {
	var c CQueue
	c.a = make([]int, 0)
	c.b = make([]int, 0)
	return c
}

func (this *CQueue) AppendTail(value int) {
	this.a = append(this.a, value)
}

func (this *CQueue) DeleteHead() (res int) {
	if len(this.b) == 0 && len(this.a) == 0 {
		return -1
	}
	if len(this.b) > 0 {
		res = this.b[len(this.b)-1]
		this.b = this.b[:len(this.b)-1]
		return
	}

	for i := len(this.a) - 1; i >= 0; i-- {
		this.b = append(this.b, this.a[i])
	}
	this.a = []int{}
	res = this.b[len(this.b)-1]
	this.b = this.b[:len(this.b)-1]

	return
}
