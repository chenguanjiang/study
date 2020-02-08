package main

import (
	"fmt"
	"reflect"
	"runtime"
)

func b1(){
	for i:=1;i<10;i++{
		for j:=1;j<=i;j++{
			fmt.Printf("%d*%d=%d ",i,j,i*j)
		}
		fmt.Println()
	}
}

func b2(a,b int)(q,r int){
	q=a/b
	r=a%b
	return
}

func b3(a, b int, method string) (int ,error) {
	switch {
	case method=="+":
		return a+b,nil
	case method=="-":
		return a-b,nil
	case method=="*":
		return a*b,nil
	case method=="/":
		return a/b,nil
	default:
		return 0,fmt.Errorf("method:'%s' not support",method)
	}
}

func apply(op func(int,int) int,a,b int) int{
	p:=reflect.ValueOf(op).Pointer()
	opname:=runtime.FuncForPC(p).Name()
	fmt.Printf("calling function %s with args (%d,%d)\n",opname,a,b)
	return op(a,b)
}

func add(a,b int) int{
	return  a*b
}

func b4(numbers ...int) int{
	sum:=0
	for i:=range numbers{
		sum+=numbers[i]
	}
	return sum
}

func swap(a,b *int){
	*a,*b=*b,*a
	fmt.Println(*a,*b)
}

func b5(){
	var arr1 [5] int=[5]int{10,2,3,4,15}
	arr2:=[3]int{1,2,3}
	arr3:=[...]int{2,4,6,8,10}
	var arr4 [4][5]int=[4][5]int{{1,2,3,4,5},{1,2,20,4,5},{1,2,3,4,5},{1,2,3,4,5}}
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr4)
	fmt.Println(arr4[1][2])
	fmt.Println("=======================")
	//for i:=0;i< len(arr1);i++  {
	//	fmt.Println(arr1[i])
	//}
	for _,v:=range arr1{
		fmt.Println(v)
	}
}

func b6(num ...int)(int,int){
	maxindex,maxvalue:=0,num[0]
	for index,value:=range num{
		if value>maxvalue{
			maxindex,maxvalue=index,value
		}
	}
	return maxindex,maxvalue
}

func b7(arr *[3]int){
	arr[0]=100
	fmt.Println(*arr)
}

func updateSlice(slice []int){
	slice[0]=100
}

func b8(){
	arr:=[...]int{0,1,2,3,4,5,6,7}
	s1:=arr[2:]
	fmt.Println("s1=",s1)
	s2:=arr[:]
	fmt.Println("s2=",s2)
	fmt.Println("update s1:=====================")
	updateSlice(s1)
	fmt.Println("s1=",s1)
	fmt.Println("arr=",arr)
	fmt.Println("update s2:=====================")
	updateSlice(s2)
	fmt.Println("s2=",s2)
	fmt.Println("arr=",arr)
	fmt.Println("################")
	s2=s2[:5]
	fmt.Println(s2)
	fmt.Println(reflect.TypeOf(s2))
	//fmt.Println(s)
	//fmt.Println(reflect.TypeOf(arr))
	//fmt.Println(reflect.TypeOf(s))
}

func main() {
	//b1()
	//q,r:=b2(13,3)
	//fmt.Println(q,r)
	//if result,err:=b3(3,4,"x");err!=nil{
	//	fmt.Printf("error:%s",err)
	//}else {
	//	fmt.Printf("result:%d",result)
	//}
	//fmt.Printf("result:%d",apply(
	//	func(a int, b int) int {
	//	return a*b
	//},
	//3,4))
	//fmt.Println(apply(add,3,4))
	//fmt.Println(b4(1,2,3,4,5))
	//a,b:=3,4
	//swap(&a,&b)
	//fmt.Println(a,b)
	//b5()
	//fmt.Println(b6(2,1,8,3,5))
	//arr:=[3]int{1,2,3}
	//b7(&arr)
	//fmt.Println(arr)
	b8()
}
