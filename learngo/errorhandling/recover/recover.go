package main

import (
	"errors"
	"fmt"
)

func tryRecover(val int) {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("Error occred:", err)
		} else {
			panic(
				fmt.Sprintf("I don't know what to do: %v", r))
		}
	}()
	if val > 0 {
		panic(errors.New("this is a error"))

	} else {
		panic(123)
	}
}

func main() {
	tryRecover(1)
	fmt.Println("tryRecover(1)")

	tryRecover(-1)
	fmt.Println("tryRecover(-1)")
}
