package main

import (
	"fmt"
	"retriever/real"
)

type Retriever interface {
	Get(url string) string
}

func download(r Retriever, url string) string {
	return r.Get(url)
}

func main() {
	var r Retriever
	r = real.RealRetriever{}
	fmt.Println(download(r, "http://www.qq.com"))
}
