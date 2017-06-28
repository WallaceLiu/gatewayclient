package main

import "fmt"

var complete chan int = make(chan int)

func loop() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}

	complete <- 0 // 执行完毕了，发个消息
}

func main() {
	go loop()
	<-complete // 直到线程跑完, 取到消息. main在此阻塞住
}
