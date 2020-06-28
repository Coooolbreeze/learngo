package fib

import (
	"fmt"
	"io"
	"strings"
)

// IntGen ...
type IntGen func() int

// Fibonacci ...
func Fibonacci() IntGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func (g IntGen) Read(p []byte) (n int, err error) {
	next := g()
	if next > 1000000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)

	// TODO: incorrect if p is too small!
	return strings.NewReader(s).Read(p)
}
