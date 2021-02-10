package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var aa = 3

var (
	ss = "kk"
	bb = true
)

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "def"
	fmt.Println(a, b, c, s)
}

func variableShorter() {
	a, b, c, s := 3, 4, true, "def"
	b = 5
	fmt.Println(a, b, c, s)
}

func euler() {
	res := cmplx.Exp(1i*math.Pi) + 1
	fmt.Println(res)
	fmt.Printf("%.2f\n", res)
}

func triangle() {
	var a, b int = 3, 4
	c := calcTriangle(a, b)
	fmt.Println(a, b, c)
}

func calcTriangle(a, b int) int {
	return int(math.Sqrt(float64(a*a + b*b)))
}

const filename = "abc.txt"

func consts() {
	const a, b = 3, 4
	c := math.Sqrt(a*a + b*b)
	fmt.Println(filename, c)
}

func enums() {
	const (
		cpp = iota
		_
		python
		golang
		javascript
	)
	fmt.Println(cpp, javascript, python, golang)

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		pb
	)
	fmt.Println(b, kb, mb, gb, pb)
}

func main() {
	fmt.Println("Hello world")

	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()

	fmt.Println(aa, ss, bb)

	euler()
	triangle()

	consts()
	enums()
}
