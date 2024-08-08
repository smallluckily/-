package main

import (
	"fmt"
	"reflect"
)

// 访问私有成员
type Person struct {
	Name string
	age  int
}

func main() {
	p := Person{
		Name: "张三",
		age:  18,
	}
	v := reflect.ValueOf(p)
	nameField := v.FieldByName("Name")
	ageField := v.FieldByName("age")
	fmt.Println(nameField.String())
	fmt.Println(ageField)
	fmt.Println(p)
}
