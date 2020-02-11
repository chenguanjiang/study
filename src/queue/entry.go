package main

import (
	"fmt"
	"queue/mypkg"
)

func main() {
	q := mypkg.Queue{1}
	q.Append(2)
	q.Append(3)
	q.Append(4)
	fmt.Println(q)
	fmt.Println("=================================")
	q.IsEmpty()
	fmt.Println("=================================")
	q.Pop()
	fmt.Println(q)
	fmt.Println("=================================")
	q.Pop()
	fmt.Println(q)
	fmt.Println("=================================")
	q.Pop()
	fmt.Println(q)
	fmt.Println("=================================")
	q.Pop()
	fmt.Println(q)
	fmt.Println("=================================")
	q.IsEmpty()
	fmt.Println("????????????????????????????????????????????")
	qq := mypkg.Queue{}
	qq.Append("a")
	qq.Append("b")
	qq.Append("c")
	qq.IsEmpty()
	fmt.Println(qq)
	qq.Pop()
	fmt.Println(qq)
	qq.Pop()
	fmt.Println(qq)
	qq.Pop()
	fmt.Println(qq)
	qq.IsEmpty()
}
