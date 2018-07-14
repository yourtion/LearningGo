package main

import (
	"fmt"
	"gointro/conc-sort/pipeline"
)

func main() {
	p := pipeline.Merge(
		pipeline.InMemSort(
			pipeline.ArraySource(3, 2, 6, 7, 4),
		),
		pipeline.InMemSort(
			pipeline.ArraySource(7, 4, 0, 3, 2, 13, 8),
		),
	)
	for v := range p {
		fmt.Println(v)
	}
}
