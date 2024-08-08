package main

import "fmt"

func main() {
	m := "qzh"
	mBytes := []byte(m)

	fmt.Println(mBytes)
	mBytes[0] = 'Q'
	m = string(mBytes)
	fmt.Println(m)
}
