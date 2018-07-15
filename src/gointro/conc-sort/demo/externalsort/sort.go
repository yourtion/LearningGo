package main

import (
	"bufio"
	"fmt"
	"gointro/conc-sort/pipeline"
	"os"
)

func main() {
	const fileIn = "large.in"
	const fileOut = "large.out"
	const n = 10000000
	const cut = 4
	genFile(os.TempDir()+fileIn, n)

	p := createPipeline(
		os.TempDir()+fileIn, n*8, cut)
	writeToFile(p, os.TempDir()+fileOut)
	printFile(os.TempDir() + fileOut)
}

func genFile(filename string, count int) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	p := pipeline.RandomSource(count)
	defer file.Close()

	writer := bufio.NewWriter(file)
	pipeline.WriterSink(writer, p)
	writer.Flush()

	fmt.Println("Gen file done")
}

func createPipeline(
	filename string,
	fileSize, chunkCount int) <-chan int {

	chunkSize := fileSize / chunkCount

	pipeline.Init()

	sortResults := []<-chan int{}
	for i := 0; i < chunkCount; i++ {
		file, err := os.Open(filename)
		if err != nil {
			panic(err)
		}

		file.Seek(int64(i*chunkSize), 0)

		source := pipeline.ReaderSource(
			bufio.NewReader(file), chunkSize)

		sortResults = append(sortResults,
			pipeline.InMemSort(source))
	}
	return pipeline.MergeN(sortResults...)
}

func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	p := pipeline.ReaderSource(file, -1)
	count := 0
	for v := range p {
		fmt.Println(v)
		count++
		if count > 100 {
			break
		}
	}
}

func writeToFile(p <-chan int, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	pipeline.WriterSink(writer, p)
}
