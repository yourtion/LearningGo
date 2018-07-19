package main

import "fmt"

func printArray(arr [5]int) {
	arr[0] = 100

	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func main() {
	var arr1 [5]int
	arr2 := [3]int{1, 3, 5}
	arr3 := [...]int{2, 4, 6, 8, 10}
	fmt.Println(arr1, arr2, arr3)

	var grid [4][5]int
	fmt.Println(grid)

	for i := 0; i < len(arr2); i++ {
		fmt.Println(arr2[i])
	}

	for i, v := range arr3 {
		fmt.Println(i, v)
	}

	fmt.Println("printArray(arr1)")
	printArray(arr1)

	fmt.Println("printArray(arr3)")
	printArray(arr3)

	fmt.Println("fmt.Println(arr1, arr3)")
	fmt.Println(arr1, arr3)
}
