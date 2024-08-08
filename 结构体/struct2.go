package main

import "fmt"

type Student3 struct {
	id      int
	name    string
	sex     byte
	age     int
	address string
}

func main() {

	var s4 Student3 = Student3{4, "yyy", 'm', 18, "shanghai"}
	var s5 Student3 = Student3{4, "yyy", 'm', 18, "shanghai"}
	fmt.Println("s4==s5", s4 == s5)
	fmt.Println("s4!=s5", s4 != s5)
}
