package main

import "fmt"

type Student struct {
	id    int
	name  string
	score int
}

func main() {
	var s1 Student = Student{1, "qzh", 100}
	fmt.Println(s1)
}
