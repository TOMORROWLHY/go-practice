package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//游戏规则简介
	var rule = `HELLO! 这是一个激烈的博弈游戏，嘿嘿嘿
游戏开始，你有10个游戏币，
每猜一次数，你将会消耗一个游戏币
每猜对一次数，你将会得到5个游戏币（是不是很丰厚？）
当你游戏币清零时，也可以向我充值，来获取更多的游戏币
1 RMB = 1 GAME_CURRENCY
只要你充的够多，将获取意想不到的惊喜哦！
`
	fmt.Println(rule)
	maxNum := 100
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(maxNum)
	// 作弊模式
	fmt.Println("The secret number is ", secretNumber)

	fmt.Println("Please input your guess")
	// 通过一个 for 循环实现一直猜数，直到猜中
	for gameCurrency := 10; ; {
		var guess int
		_, err := fmt.Scanf("%d", &guess)
		if err != nil {
			fmt.Println("Invalid input. Please enter an integer value")
			continue
		}
		fmt.Println("You guess is", guess)
		if guess > secretNumber {
			fmt.Println("Your guess is bigger than the secret number. Please try again")
			gameCurrency = gameCurrency - 1
			fmt.Println("your GAME_CURRENCY :", gameCurrency)
		} else if guess < secretNumber {
			fmt.Println("Your guess is smaller than the secret number. Please try again")
			gameCurrency = gameCurrency - 1
			fmt.Println("your GAME_CURRENCY :", gameCurrency)
		} else {
			fmt.Println("Correct, you Legend!")
			gameCurrency = gameCurrency + 5
			fmt.Println("your GAME_CURRENCY :", gameCurrency)
			break
		}
	}
}
