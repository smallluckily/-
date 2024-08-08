package main

import "fmt"

type Student struct {
	id   int
	name string
}

func (s Student) PrintIofo() {
	fmt.Println(s.id, s.name)
}

func main() {
	s := Student{1, "qzh"}
	s.PrintIofo()
}
