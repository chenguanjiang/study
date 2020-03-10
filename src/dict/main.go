package main

import (
	. "dict/myinterface"
	"dict/mystruct"
	"fmt"
)

func put(w WriterTool, i interface{}) {
	w.Write(i)
}

func get(r ReaderTool, index int) interface{} {
	return r.Read(index)
}

func main() {
	ms := &mystruct.MyStruct{[]interface{}{}}
	var w WriterTool
	var r ReaderTool
	w = ms
	r = ms
	put(w, "abc")
	put(w, 1)
	put(w, map[string]string{"name": "gjchen"})
	fmt.Println(w)
	fmt.Println(get(r, 0))
	fmt.Println(get(r, 1))
	fmt.Println(get(r, 2))
	fmt.Println("-------------------------------------")
	fmt.Println(r)
}
