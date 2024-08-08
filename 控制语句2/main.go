package main

import "fmt"

func main() {
	var date int
	fmt.Println("请输入你的出生日期用来判断你的星座：")
	fmt.Scan(&date)
	switch date {
	case 1, 2, 3:
		fmt.Println("你是双子座")
	case 4, 5, 6:
		fmt.Println("你是双鱼座")
	default:
		fmt.Println("输入错误")

	}
}
