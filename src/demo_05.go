package main

import (
	"fmt"
	"time"
)

type Person struct {
	name string
	age  int
}

type Class struct {
	num   int
	price int
}

type Sa struct {
	Person
	Class
	flag bool
}

func (cl *Class) Say(s string) {
	fmt.Println(cl.num, ",price:", cl.price, "string:", s)
}

type NUM int

func say_hello01(n NUM) {
	fmt.Println(n)
}

type My01 struct {
	name string
	age  int
}

func (m *My01) String() string {
	return fmt.Sprintf("name:%s,age:%d", m.name, m.age)
}

func gjchen_01(a, b int) (new_a, new_b int) {
	a, b = b, a
	return a, b
}

//func gjchen_02(l []int) []int {
//	for index, _ := range l {
//		for ind, _ := range l {
//			if ind < len(l)-1 {
//				if l[ind] > l[ind+1] {
//					l[ind], l[ind+1] = l[ind+1], l[ind]
//				}
//			}
//		}
//	}
//	return l
//}

func main() {
	//say_hello01(0)
	//m:=&My01{"gjchen",20}
	//fmt.Println(m)
	//fmt.Println(gjchen_01(1,2))
	fmt.Println("##########:", time.Second)
}
