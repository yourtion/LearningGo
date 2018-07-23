package main

import (
	"fmt"
	"retriever/mock"
	"retriever/real"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("https://blog.yourtion.com")
}

func main() {
	var r Retriever

	r = mock.Retriever{"this is a mock"}
	fmt.Println(download(r))

	r = real.Retriever{}
	fmt.Println(download(r))
}
