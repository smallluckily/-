package main

import "fmt"

// 人
type Person1 struct {
	name string
	sex  byte
	age  int
}

// 学生
type Student5 struct {
	Person1 //匿名字段，那么默认Student就包含了Person的所有字段
	id      int
	address string
}

func main() {
	var s1 Student5 //变量声明
	//给成员赋值
	s1.name = "mike"
	s1.sex = 'm'
	s1.age = 18
	s1.id = 1
	s1.address = "shanghai"
	fmt.Println(s1) //{{mike 109 18} 1 shanghai}

	var s2 Student5 //变量声明
	s2.Person1 = Person1{"Lily", 'f', 19}
	s2.id = 2
	s2.address = "beijing"
	fmt.Println(s2) //	{{Lily 102 19} 2 beijing}
}
