package mypkg

import "fmt"

type Queue []int

func (q *Queue) Append(num int){
	*q=append(*q,num)
}

func (q *Queue) Pop(){
	*q=(*q)[1:]
}

func (q *Queue) IsEmpty(){
	fmt.Println(len(*q)==0)
}
