package main

import (
	"fmt"
	"gointro/conc-sort/pipeline"
)

func main() {
	p := pipeline.ArraySource(3, 2, 6, 7, 4)
	for v := range p {
		fmt.Println(v)
	}
}
