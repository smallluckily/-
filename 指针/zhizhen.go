package main

import "fmt"

func main() {
	var a int = 10            //声明一个变量，同时初识化
	fmt.Printf("&a=%p\n", &a) //操作符“&” 取变量地址
	var p *int = nil          //声明一个变量p,类型为*int,指针类型
	p = &a
	fmt.Printf("p=%p\n", p)
	fmt.Printf("a=%d,*p=%d\n", a, *p)
}
