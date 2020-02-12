package mystruct

type MyStruct struct {
	Content []interface{}
}

func (ms *MyStruct) Write(thing interface{}) {
	ms.Content = append(ms.Content, thing)
}

func (ms *MyStruct) Read(index int) interface{} {
	return ms.Content[index]
}
