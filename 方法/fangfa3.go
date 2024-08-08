package main

import "fmt"

type Person1 struct {
	name string
	sex  byte
	age  int
}

//Person定义了方法

func (p *Person1) PrintInfo() { //给Person添加方法
	fmt.Println(p.name, p.sex, p.age)
}

type Student1 struct {
	Person1 //匿名字段，那么Student包含了Person的所有字段
	id      int
	address string
}

func main() {
	p := Person1{"mike", 'm', 18}
	p.PrintInfo()
	s := Student1{Person1{"Lily", 'f', 20}, 2, "shanghai"}
	s.PrintInfo()
}
