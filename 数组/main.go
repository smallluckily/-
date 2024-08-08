package main

import "fmt"

func main() {
	var a = [...]int{1, 2, 3, 4}
	slice2 := a[:]
	fmt.Println(slice2)
	slice1 := a[0:3]
	fmt.Println(slice1)
	slice1[2] = 0
	fmt.Println(slice1)
	fmt.Println(a)
	fmt.Println(slice2)
}
