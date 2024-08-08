package main

import "fmt"

// 使用值为nil的map会发生什么？使用值为nil的slice会发生什么？
// 答：允许对值为nil的slice进行append操作添加元素，但对值为nil的map进行操作会引发panic。
func main() {
	/*var m map[string]int
	m["one"] = 1
	fmt.Printf("m")*/

	m := map[string]string{"one": "1", "two": "", "three": "3"}
	if _, ok := m["two"]; !ok {
		fmt.Println("two is not set")
	}

	x := map[string]string{"one": "2", "two": "2", "three": "3"}
	if _, ok := x["two"]; !ok {
		fmt.Println("key two is no entry")
	}

}
