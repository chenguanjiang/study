package both

type BothData struct {
	some string
	any  string
}

func (b *BothData) Post(url string) string {
	return b.some + b.any
}
