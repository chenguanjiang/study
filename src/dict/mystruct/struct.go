package mystruct

import "fmt"

type MyStruct struct {
	Content []interface{}
}

func (ms *MyStruct) Write(thing interface{}) {
	ms.Content = append(ms.Content, thing)
}

func (ms *MyStruct) Read(index int) interface{} {
	return ms.Content[index]
}

func (ms *MyStruct) String() string {
	return fmt.Sprintf("Content:%v", ms.Content)
}
