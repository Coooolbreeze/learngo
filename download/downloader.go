package main

import (
	"fmt"
	"learngo/download/infra"
)

func getRetriever() retriever {
	return infra.Retriever{}
}

type retriever interface {
	Get(string) string
}

func main() {
	retriever := getRetriever()

	fmt.Println(retriever.Get("https://www.imooc.com"))
}
