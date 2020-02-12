package myinterface

type WriterTool interface {
	Write(thing interface{})
}

type ReaderTool interface {
	Read(index int) interface{}
}
