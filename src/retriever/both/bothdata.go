package both

type BothData struct {
	Some  string
	Any   string
	Thing string
}

type Love struct {
}

func (b *BothData) Post(url string) string {
	return b.Some + b.Any + url
}

func (b *BothData) Get(url string) string {
	return b.Thing + url
}

func (b *BothData) Love(url string) string {
	return "my love"
}

func (l *Love) Love(url string) string {
	return "hello love"
}
