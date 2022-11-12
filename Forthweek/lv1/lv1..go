package main

import (
	"fmt"
)

func main() {
	var n int
	var count []int
	var numCount = make(map[int]int)
	fmt.Println("你要输入几行自然数？")
	fmt.Scan(&n)
	//向map中输入元素
	for i := 0; i < n; i++ {
		var a int
		fmt.Scan(&a)
		count = append(count, a)
		numCount[a] = 0
	}
	fmt.Println(count)
	//对map中的数进行统计
	for _, num := range count {
		v, ok := numCount[num]
		if ok {
			numCount[num] = v + 1
		}
	}
	for k, v := range numCount {
		fmt.Println(k, v)
	}
}
