package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p *Person) SetInfoPointer() { //指针作为接收者，引用语义
	//给成员赋值
	p.name = "Lily"
	p.age = 24
	fmt.Println(*p)
}
func main() {
	p := Person{"qzh", 22}
	p.SetInfoPointer()
	fmt.Println(p)
}
