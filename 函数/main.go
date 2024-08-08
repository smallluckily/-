package main

import "fmt"

func calc(x, y int) (sum, sub int) {
	sum = x + y
	sub = x - y
	return //  return sum sub
}
func main() {
	fmt.Println(calc(3, 3))
}
