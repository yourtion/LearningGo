package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(
				time.Duration(rand.Intn(1500)) *
					time.Millisecond)
			out <- i
			i++
		}

	}()
	return out
}

func worker(id int, c <-chan int) {
	for n := range c {
		time.Sleep(time.Second)
		fmt.Printf("Worker %d recevved %d \n", id, n)
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func main() {
	var c1, c2 = generator(), generator()
	worker := createWorker(0)

	// 缓冲 value
	var values []int

	tm := time.After(10 * time.Second)
	tick := time.Tick(time.Second)

	for {
		var activeWorker chan<- int
		var activeValues int
		if len(values) > 0 {
			activeWorker = worker
			activeValues = values[0]
		}

		select {

		case n := <-c1:
			values = append(values, n)
		case n := <-c2:
			values = append(values, n)
		case activeWorker <- activeValues:
			values = values[1:]

		case <-time.After(800 * time.Millisecond):
			// no data in 800ms
			fmt.Println("timeout")
		case <-tick:
			// report queue length every 1s
			fmt.Println("queue len = ", len(values))
		case <-tm:
			// run 10s and exit
			fmt.Println("Bye.")
			return
		}
	}
}
