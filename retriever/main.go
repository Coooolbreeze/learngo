package main

import (
	"fmt"
	"learngo/retriever/mock"
	"learngo/retriever/real"
	"time"
)

// Retriever ...
type Retriever interface {
	Get(url string) string
}

// Poster ...
type Poster interface {
	Post(url string, form map[string]string) string
}

const url = "https://www.imooc.com"

func download(r Retriever) string {
	return r.Get(url)
}

func post(poster Poster) string {
	return poster.Post(url, map[string]string{
		"name":   "liwenxiang",
		"course": "golang",
	})
}

// RetrieverPoster ...
type RetrieverPoster interface {
	Retriever
	Poster
}

func session(s RetrieverPoster) string {
	s.Post(url, map[string]string{
		"contents": "Another faked imooc.com",
	})
	return s.Get(url)
}

func inspect(r Retriever) {
	fmt.Println("Inspecting", r)
	fmt.Printf(" > %T %v\n", r, r)

	switch v := r.(type) {
	case *mock.Retriever:
		fmt.Println("Contents:", v.Contents)
	case *real.Retriever:
		fmt.Println("UserAgent:", v.UserAgent)
	}

	fmt.Println()
}

func main() {
	retriever := mock.Retriever{
		Contents: "This is a fake imooc.com",
	}

	var r Retriever
	r = &mock.Retriever{Contents: "This is a fake imooc.com"}
	inspect(r)

	r = &real.Retriever{
		UserAgent: "Mozilla/5.0",
		Timeout:   time.Minute,
	}
	inspect(r)

	if realRetriever, ok := r.(*real.Retriever); ok {
		fmt.Println(realRetriever.Timeout)
	} else {
		fmt.Println("Not a real retriever.")
	}

	if mockRetriever, ok := r.(*mock.Retriever); ok {
		fmt.Println(mockRetriever.Contents)
	} else {
		fmt.Println("Not a mock retriever.")
	}

	// fmt.Println(download(r))

	fmt.Println(session(&retriever))

}
