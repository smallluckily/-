package main

import "fmt"

func main() {
	m1 := map[int]string{0: "qzh", 1: "mc"}
	fmt.Println(m1)
	m1[1] = "zj"
	m1[3] = "z"
	fmt.Println(m1)
	m2 := m1
	fmt.Println(m2)
	m2[1] = "qzh"
	fmt.Println(m2)
	fmt.Println(m1)
}
