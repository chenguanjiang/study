package main

import (
	"dict/myinterface"
	"dict/mystruct"
	"fmt"
)

func put(w myinterface.WriterTool, i interface{}) {
	w.Write(i)
}

func get(r myinterface.ReaderTool, index int) interface{} {
	return r.Read(index)
}

func main() {
	ms := &mystruct.MyStruct{[]interface{}{}}
	var w myinterface.WriterTool
	var r myinterface.ReaderTool
	w = ms
	r = ms
	put(w, "abc")
	put(w, 1)
	put(w, map[string]string{"name": "gjchen"})
	fmt.Println(w)
	fmt.Println(get(r, 0))
	fmt.Println(get(r, 1))
	fmt.Println(get(r, 2))
}
