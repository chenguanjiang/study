package main

import (
	"fmt"
	"reflect"
	"unicode/utf8"
)

func c1() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
	s1 := arr[2:6]
	s2 := s1[3:5]
	fmt.Println(s1)
	fmt.Println(s2)
	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s4 = append(s4, 12)
	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Println(arr)
}

func c2() {
	var s []int
	fmt.Println(s)
	for i := 0; i < 100; i++ {
		s = append(s, 2*i+1)
		print_slice(s)
	}
	fmt.Println(s)
}

func print_slice(s []int) {
	fmt.Printf("slice len:%d,cap:%d\n", len(s), cap(s))
}

func c3() {
	s1 := []int{2, 4, 6, 8}
	fmt.Println(s1)
	s2 := make([]int, 10)
	s3 := make([]int, 10, 12)
	var s4 []int
	s4 = []int{1, 2, 3, 4}
	//s4:=[]int{}
	fmt.Println(s2)
	fmt.Println(s3)
	print_slice(s2)
	print_slice(s3)
	//fmt.Printf("%v",s2)
	//fmt.Printf("%v",s3)
	fmt.Println(s4)
	fmt.Println("=================================")
	copy(s2, s4)
	fmt.Println(s2)
	print_slice(s2)
	s2 = append(s2, 10)
	fmt.Println(s2)

}

func cc3() {
	s := []int{2, 4, 6, 8, 10}
	for i := range s {
		fmt.Println(i)
	}
}

func c4(s []int, i int) {
	fmt.Println(s)
	new_s := append(s[:i], s[i+1:]...)
	fmt.Println(new_s)
	fmt.Println(s)
}

func c5() {
	m1 := map[string]string{
		"gjchen": "father",
		"yan":    "mother",
		"junhan": "son",
	}
	m2 := make(map[string]int)
	var m3 map[string]int = map[string]int{
		"d": 1,
		"e": 2,
		"f": 3,
	}
	fmt.Println(m1)
	fmt.Println(m2)
	fmt.Println(m3)
	fmt.Println("======================")
	for k, v := range m1 {
		fmt.Println(k, v)
	}
	fmt.Println("======================")
	fmt.Println(m1["gjchen"])
	fmt.Println("======================")
	fmt.Println(reflect.TypeOf(m3["v"]))
	fmt.Println("======================")
	//value,ok:=m1["gjchen"]
	//fmt.Println(value,ok)
	//if ok  {
	//	fmt.Println("yes")
	//}else {
	//	fmt.Println("no")
	//}
	if value, ok := m1["gjchenn"]; ok {
		fmt.Println(value)
	} else {
		fmt.Println("key does not exit")
	}
	fmt.Println("======================")
	fmt.Println(m1)
	if _, ok := m1["gjchen"]; ok {
		delete(m1, "gjchen")
	}
	fmt.Println(m1)
}

func c6() {
	s := "冠江abcde"
	//for i:=0;i<len(s);i++{
	//	fmt.Println(s[i])
	//}
	//for _,e:=range s{
	//	fmt.Printf("%c",e)
	//}
	//n_s:=[]byte(s)
	//fmt.Println(n_s)
	//for _,e:=range n_s{
	//	fmt.Printf("%c",e)
	//}
	var runes = []rune(s)
	fmt.Println(runes)
	for _, e := range runes {
		fmt.Printf("%c", e)
	}
}

func GetNoRepeatStringMaxLength(str string) int {
	start := 0
	maxlength := 0
	str_index_map := make(map[rune]int)
	for index, key := range []rune(str) {
		if lastIndex, ok := str_index_map[key]; ok && lastIndex >= start {
			start = str_index_map[key] + 1
		}
		if index-start+1 > maxlength {
			maxlength = index - start + 1
		}
		str_index_map[key] = index
	}

	return maxlength
}

func c7() {
	s := "abc我爱中国！"
	n := utf8.RuneCountInString(s)
	fmt.Println(n)
	fmt.Println(len(s))
	for i, v := range s {
		fmt.Println(i, v)
	}
	for i, v := range []byte(s) {
		fmt.Println(i, v)
	}
}

func c8() {
	type treeNode struct {
		value       int
		left, right *treeNode
	}
	var root treeNode
	fmt.Println(root)
	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode)
	fmt.Println(root)
	r := treeNode{3, nil, nil}
	fmt.Println(r)
	fmt.Println("=====================================")
	root1 := treeNode{}
	root1.value = 5
	root1.left = &treeNode{}
	fmt.Println(root1)
	fmt.Println("=====================================")
	nodes := []treeNode{
		{3, nil, nil},
		{},
		{value: 3},
		{left: nil},
		{right: &treeNode{}},
	}
	fmt.Println(nodes)
}

type nodeTree struct {
	value int
	left  *nodeTree
	right *nodeTree
}

func c9(value int) *nodeTree {
	return &nodeTree{value: value}
}

func main() {
	//c4([]int{1,2,3,4},0)
	//c6()
	//str := "abcabcbb我爱中国呀哈哈哈哈哈！"
	//maxLength := GetNoRepeatStringMaxLength(str)
	//fmt.Println(maxLength)
	fmt.Println(c9(3))
}
