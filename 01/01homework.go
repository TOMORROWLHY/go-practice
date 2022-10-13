package main

import (
	"fmt"
	"math"
)

var (
	a = 0
	b = 55
	c = math.MaxInt //改成10000000000运行是没有问题的，这个maxint，我根本不敢运行，我怕电脑炸裂
	d = 0
)

func main() {

	for i := 0; a < c; i = i + 1 {
		d = int(math.Pow(10, float64(i)))
		a = d*(b-5) + 5
		fmt.Println(a)
	}
}
