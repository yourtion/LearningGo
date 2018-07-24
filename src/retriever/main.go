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

type Poster interface {
	Post(url string,
		form map[string]string) string
}

const url = "https://blog.yourtion.com"

func download(r Retriever) string {
	return r.Get(url)
}

func post(poster Poster) {
	poster.Post(url, map[string]string{
		"name":   "Yourtion",
		"course": "golang",
	})
}

type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents": "another fake blog",
	})
	return s.Get(url)
}

func main() {
	var r Retriever

	retriever := &mock.Retriever{"this is a mock"}
	inspect(retriever)

	fmt.Println(download(retriever))
	post(retriever)

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
	fmt.Println("Try a session")
	fmt.Println(session(retriever))
}

func inspect(r Retriever) {
	fmt.Printf(" > %T, %v \n", r, r)
	fmt.Print(" > Type Switch: ")
	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}
	fmt.Println("Type switch done")
}
