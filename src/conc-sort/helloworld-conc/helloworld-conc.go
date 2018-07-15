package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	for i := 0; i < 5000; i++ {
		 // go printHelloWorld(i)
		 go printHelloWorld2(i, ch)
	}
	for {
		msg := <- ch
		fmt.Println(msg)
	}
	time.Sleep(time.Millisecond)
}

func printHelloWorld(i int) {
	for {
		fmt.Printf("Hello World %d n", i)
	}
}

func printHelloWorld2(i int, ch chan string) {
	for {
		ch <- fmt.Sprintf("Hello World %d !\n", i)
	}
}
