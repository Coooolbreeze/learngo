package real

import (
	"net/http"
	"net/http/httputil"
	"time"
)

// Retriever ...
type Retriever struct {
	UserAgent string
	Timeout   time.Duration
}

// Get ...
func (r *Retriever) Get(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	result, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	resp.Body.Close()

	return string(result)
}
