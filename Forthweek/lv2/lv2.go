package main

import "fmt"

func main() {
	fmt.Println("提醒功能如下：")
	fmt.Println("1.单次日程提醒")
	fmt.Println("2.重复日程提醒")
	fmt.Println("请输入序号")
	var input int
	fmt.Scan(&input)
	switch input {
	case 1:
	case 2:
	default:
		fmt.Println("无效的数字")
	}
}
func setTime_once() {
	fmt.Println("请设置提醒时间（格式为：2006-01-02 15:04:05）：")

}
