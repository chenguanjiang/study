package real

import (
	"net/http"
	"net/http/httputil"
	"time"
)

type RealRetriever struct {
	UserAgent string
	TimeOut   time.Duration
}

func (r RealRetriever) Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	result, err := httputil.DumpResponse(resp, true)
	resperr := resp.Body.Close()
	if resperr != nil {
		panic(resperr)
	}
	if err != nil {
		panic(err)
	}
	return string(result)
}
