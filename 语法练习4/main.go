package main

func main() {
	for i := 0; i < 5; i++ {
		//i := i //删除这行就只是输出5，现在是输出4，3，2，1，0
		defer func() {
			println(i)
		}()
	}
}
