package main

import (
	"fmt"
	"retriever/mock"
	"retriever/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("https://blog.yourtion.com")
}

func main() {
	var r Retriever

	// 值类型的接口实现可以指针传递
	r = &mock.Retriever{"this is a mock"}
	inspect(r)

	fmt.Println(download(r))

	r = &real.Retriever{
		UserAgent: "Yourtion-Go",
		TimeOut:   time.Minute,
	}
	inspect(r)

	// Type assertion
	realRetriever := r.(*real.Retriever)
	fmt.Println(realRetriever.TimeOut)

	// Type assertion error
	if mockRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("not a mock retriever")
	}

	//fmt.Println(download(r))
}

func inspect(r Retriever) {
	fmt.Printf("%T, %v \n", r, r)
	fmt.Println("Type Switch")
	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
	fmt.Println("Type switch done")
}
