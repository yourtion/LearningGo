package main

import (
	"fmt"
	"time"
)

func worker(id int, c <-chan int) {
	for n := range c {
		// n, ok := <-c
		// if !ok {
		//	 break
		// }
		fmt.Printf("Worker %d recevved %c \n", id, n)
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func chanDemo() {
	var channels [10]chan<- int

	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'A' + i
	}

	time.Sleep(time.Millisecond)

}

func bufferedChannel() {
	c := make(chan int, 3)
	go worker(10, c)

	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'

	time.Sleep(time.Millisecond)
}

func channelClose() {
	c := make(chan int)
	go worker(11, c)

	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'

	close(c)
	time.Sleep(time.Millisecond)
}

func main() {
	fmt.Println("Channel as first-class citizen")
	chanDemo()

	fmt.Println("Buffered Channel")
	bufferedChannel()

	fmt.Println("Channel close and range")
	channelClose()
}
