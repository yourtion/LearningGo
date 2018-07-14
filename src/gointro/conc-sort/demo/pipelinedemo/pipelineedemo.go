package main

import (
	"bufio"
	"fmt"
	"gointro/conc-sort/pipeline"
	"os"
)

func main() {
	const filename = "large.in"
	const n = 10000000
	file, err := os.Create(os.TempDir() + filename)
	if err != nil {
		panic(err)
	}
	p := pipeline.RandomSource(n)
	defer file.Close()

	writer := bufio.NewWriter(file)
	pipeline.WriterSink(writer, p)
	writer.Flush()

	file, err = os.Open(os.TempDir() + filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p = pipeline.ReaderSource(
		bufio.NewReader(file))
	count := 0

	for v := range p {
		fmt.Println(v)
		count++
		if count >= 100 {
			break
		}
	}
}
