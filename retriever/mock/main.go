package mock

import "fmt"

// Retriever ...
type Retriever struct {
	Contents string
}

func (r *Retriever) String() string {
	return fmt.Sprintf("Retriever: {Contents=%s}", r.Contents)
}

// Get ...
func (r *Retriever) Get(url string) string {
	return r.Contents
}

// Post ...
func (r *Retriever) Post(url string, form map[string]string) string {
	r.Contents = form["contents"]
	return "ok"
}
