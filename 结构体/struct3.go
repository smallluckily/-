package main

import "fmt"

type Student4 struct {
	id      int
	name    string
	sex     byte
	age     int
	address string
}

func printStudentValue(tmp Student4) {
	tmp.id = 250
	fmt.Println("printStudentValue tmp=", tmp)
}

func main() {
	var s4 Student4 = Student4{4, "yyy", 'm', 18, "shanghai"}
	printStudentValue(s4) //值传递,形参修改不影响实参
	fmt.Println("main s4=", s4)
}
