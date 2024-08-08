package main

import "fmt"

type Student2 struct {
	id      int
	name    string
	sex     byte
	age     int
	address string
}

func main() {
	s1 := new(Student2)
	s1.id = 1
	s1.name = "qzh"
	fmt.Println(s1, *s1)

	var s4 Student2 = Student2{4, "yyy", 'm', 18, "shanghai"}
	fmt.Println(s4, &s4)

	var p *Student2 = &s4
	p.id = 5
	(*p).name = "zzz"
	fmt.Println(p, *p, s4)

}
