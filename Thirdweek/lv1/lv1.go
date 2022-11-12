package main

import (
	"fmt"
	"time"
)

//自己想的不知道行不行
//func main() {
//好康的 := 欢迎来我家玩()
//fmt.Println(好康的)
//}

//func 欢迎来我家玩() string {
// 阿伟想在去杰哥家的路上打电动
//打电动()
// 花费 5s 前往杰哥家
//time.Sleep(5 * time.Second)
//return "登dua郎"
//}

//func 打电动() {
//fmt.Println("输了啦，都是你害的")

//}

// 回调

func main() {
	surprise()

}
func welcome() string {
	//花费五秒到达杰哥家
	time.Sleep(5 * time.Second)
	return "登dua郎"
}
func playGames() {
	fmt.Println("垃圾队友！输啦！")
}
func surprise() {
	playGames()
	fmt.Println("死啦！都是你害的！")
	fmt.Println(welcome())
}
