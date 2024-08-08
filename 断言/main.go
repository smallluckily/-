package main

import "fmt"

// 不会拷贝数据，只是改变类型
func main() {
	var m interface{} = 1
	if v, ok := m.(int); !ok {
		fmt.Println("类型断言失败")
	} else {
		fmt.Println("类型断言成功", v)
	}
}
