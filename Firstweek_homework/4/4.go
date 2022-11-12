package main

import "fmt"

func main() {
	fmt.Println(0)
	fmt.Println(10)
	for i := 11; i < 100000; i++ {
		var a string = string(i)
		length := len(a)
		for j := 1; j < length; j++ {
			b := a[0]
			c := a[j]
			if b+c == 10 {
				fmt.Println(a)
			}

		}

	}
}
