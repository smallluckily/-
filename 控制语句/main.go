package main

import "fmt"

func main() {
	var age int
	fmt.Println("请输入你的年龄：")
	fmt.Scan(&age)
	if age >= 18 {
		fmt.Println("你是成年人可以进行此操作")
	} else {
		fmt.Println("你还未成年不能进行这个操作")
	}
}
