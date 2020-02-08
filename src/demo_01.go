package main

import (
	"fmt"
	"math"
	"strconv"
)

func myfun1() int {
	var a,b,c int=1,2,3
	return a+b+c
}

func myfun2() string{
	a:=3
	b:="abc"
	c:=strconv.Itoa(a)
	d:=b+c
	return d
}

func myfun3(){
	a,b,c:=1,true,"abc"
	fmt.Println(a,b,c)
}

//var a1,a2=1,2

func myfun4(){
	var(
		a=1
		b=2
		c=3
	)
	fmt.Println(a,b,c)
}

func myfun5(){
	var(
		a int
		b string
		c int
	)
	a,b,c=1,"aaa",2
	fmt.Printf("type:%T,value:%v\n",a,a)
	fmt.Printf("type:%T,value:%v\n",b,b)
	fmt.Printf("type:%T,value:%v\n",c,c)
}

func myfun6(){
	var a,b int=3,4
	var c float64
	c=math.Sqrt(float64(a*a+b*b))
	fmt.Println(c)
}

func myfun7(){
	var a,b=3,4
	c:=math.Sqrt(a*a+b*b)
	fmt.Println(c)
}

func main(){
	//fmt.Println(myfun2())
	//myfun3()
	//fmt.Println(a1,a2)
	//myfun4()
	myfun7()
}




