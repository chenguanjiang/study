package main

import (
	"fmt"
	"retriever/both"
	"retriever/mock"
	"retriever/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string) string
}

type LoveI interface {
	Love(string) string
}

type RetrieverPoster interface {
	Retriever
	Poster
	Love(url string) string
}

func download(r Retriever, url string) string {
	return r.Get(url)
}

func post(p Poster, url string) string {
	return p.Post(url)
}

func download_post(dp RetrieverPoster, url string) string {
	return dp.Post(url)
}

func get_love(rpl RetrieverPoster, url string) string {
	return rpl.Love(url)
}

func get_hello_love(l LoveI, url string) string {
	return l.Love(url)
}

func inspect(r Retriever) {
	fmt.Printf("%T,%v \n", r, r)
	switch v := r.(type) {
	case mock.MockRetriever:
		fmt.Println("Contents:", v.Contents)
	case *real.RealRetriever:
		fmt.Println("UserAgent", v.UserAgent)
	}
}

func main() {
	//fmt.Println(download(r, "http://www.qq.com"))

	var r Retriever
	//fmt.Printf("%T,%v",r,r)
	r = mock.MockRetriever{"hello world"}
	inspect(r)
	r = &real.RealRetriever{
		UserAgent: "Mozilla/5.0",
		TimeOut:   time.Minute,
	}
	inspect(r)
	fmt.Println("===========================")
	if f, ok := r.(*real.RealRetriever); ok {
		fmt.Println(f.TimeOut)
	} else {
		fmt.Println("r is not a RealRetriever type")
	}
	fmt.Println("===========================")
	fmt.Println("#########################################")
	var rp RetrieverPoster
	rp = &both.BothData{"hello", "world", "gjchen"}
	fmt.Println(download(rp, "!"))
	fmt.Println(post(rp, "!"))
	fmt.Println(download_post(rp, "!"))
	var l LoveI
	l = &both.Love{}
	fmt.Println(get_hello_love(l, ""))
}
