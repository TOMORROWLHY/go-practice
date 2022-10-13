package main

import (
	"fmt"
)

func main() {
	var name1 string
	fmt.Println("请输入小组名称")
	var name2 string
	fmt.Println("请输入与UFO名称")
	fmt.Scanf("%s", &name1)
	fmt.Scanf("%s", &name2)
	var (
		len1 = len(name1)
		len2 = len(name2)
		a1   = 1
		a2   = 1
	)

	for i := 0; i < len1; i++ {
		a1 = (int(name1[i]) - 64) * a1
	}
	for i := 0; i < len2; i++ {
		a2 = (int(name2[i]) - 64) * a2
	}
	fmt.Println(a1) //查看乘积是否正确
	fmt.Println(a2) //查看乘积是否正确
	if a1%47 == a2%47 {
		fmt.Println("GO GO GO !!!")
	} else {
		fmt.Println("Do you want to go? No way!!!")
	}
}
