package mock

type MockRetriever struct {
	Contents string
}

func (r MockRetriever) Get(url string) string {
	return r.Contents
}
