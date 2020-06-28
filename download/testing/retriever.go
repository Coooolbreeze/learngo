package testing

type Retriever struct{}

// Get ...
func (Retriever) Get(url string) string {
	return "Fake content"
}
