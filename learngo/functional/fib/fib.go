package fib

import (
	"fmt"
	"io"
	"strings"
)

// 1, 1, 2, 3, 5, 8, 13, 21
//    a, b
//       a, b
func Fibonacci() IntGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

type IntGen func() int

func (g IntGen) Read(p []byte) (n int, err error) {
	next := g()
	if next > 10000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)

	// TODO: incorrect if p is oo small!
	return strings.NewReader(s).Read(p)
}
