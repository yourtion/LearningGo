package main

import (
	"bufio"
	"fmt"
	"learngo/functional/fib"
	"os"
)

func tryDefer() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	//panic("error occurred")
	return
	fmt.Println(4)
}

func tryDefer2() {
	for i := 0; i < 100; i++ {
		defer fmt.Println("tryDefer2: ", i)
		if i == 5 {
			//panic("error")
			return
		}
	}
}

func writeFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fib.Fibonacci()
	for i := 0; i < 20; i++ {
		fmt.Fprintln(writer, f())
	}
}

func main() {
	writeFile("fib.out")

	tryDefer()
	tryDefer2()
}
