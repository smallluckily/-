package main

import "fmt"

func main() {
	/*a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	b := a
	fmt.Println(a, b)
	b[0] = 0
	fmt.Println(b)
	fmt.Println(a)*/

	//var a = []int{1, 2, 3, 4, 5}
	//b := append(a, 44)
	//c := a
	//fmt.Println(a, b)
	//fmt.Println(c)
	//a[0] = 0
	//fmt.Println(a, b)
	//fmt.Println(c)

	var a = [4]int{1, 2, 3, 4}
	s1 := a[0:2] //{1,2}
	s2 := s1     //{1,2}
	s3 := append(append(append(s1, 100), 200), 300)
	s1[0] = 1000
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(a)
}
