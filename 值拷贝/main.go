package main

import (
	"fmt"
	"strings"
)

func main() {
	var a = 10
	var b = 20
	a = 30
	fmt.Println(a)
	fmt.Println(b)

	var d = 100
	var e = 200
	d, e = e, d
	fmt.Println(d, e)

	s := "我的 你的  他的"
	slice := strings.Split(s, " ")
	fmt.Println(slice)
	slice1 := strings.Join(slice, " ,")
	fmt.Println(slice1)

	index1 := strings.Index(slice1, "我的")
	fmt.Println(index1)
	name := "yuan"
	age := 24
	isMarried := false
	salary := 3000.549
	fmt.Printf("姓名:%s 年龄:%d 婚否：%t 薪资:%.2f\n", name, age, isMarried, salary)
	fmt.Printf("姓名:%v 年龄:%v 婚否：%v 薪资:%v\n", name, age, isMarried, salary)
	fmt.Printf("姓名:%#v 年龄:%#v 婚否：%#v 薪资:%#v\n", name, age, isMarried, salary)


