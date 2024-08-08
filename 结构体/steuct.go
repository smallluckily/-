package main

import "fmt"

type Student1 struct {
	id      int
	name    string
	sex     byte
	age     int
	address string
}

func main() {
	var s1 Student1 = Student1{1, "", 'm', 18, "shanghai"}
	fmt.Printf("id=%d,name=%s,sex=%c,age=%d,address=%s\n", s1.id, s1.name, s1.sex, s1.age, s1.address)

	var s2 Student1
	s2.id = 2
	s2.name = "lily"
	s2.sex = 'm'
	s2.age = 18
	s2.address = "shanghai"
	fmt.Println(s2)
}
