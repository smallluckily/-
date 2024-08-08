package main

import "fmt"

func main() {
	m := 10
	s := "qzh"
	fmt.Println(m, s)
	func() {
		m = 20
		s = "zhang"
		fmt.Println(m, s)
	}()
	fmt.Println(m, s)
}
