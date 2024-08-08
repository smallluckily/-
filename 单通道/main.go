package main

import (
	"fmt"
)

func counter(out chan<- int) { // chan<-只写
	defer close(out)
	for i := 0; i < 5; i++ {
		out <- i //如果对方不读，会阻塞
	}
}

func printer(in <-chan int) { // <-chan只读
	for num := range in {
		fmt.Println(num)
	}
}

func main() {
	c := make(chan int) //chan 读写

	//生产者，消费者
	go counter(c)
	printer(c)
	fmt.Println("done")
}
