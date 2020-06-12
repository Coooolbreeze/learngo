package main

import (
	"fmt"
	"math/cmplx"
)

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func euler() {
	c := 3 + 4i
	fmt.Println(cmplx.Abs(c))
}

func enums() {
	const (
		cpp = iota
		java
		python
		golang
	)

	fmt.Println(cpp, java, python, golang)
}

func main() {
	fmt.Println("Hello World!")
	variableZeroValue()
	euler()
	enums()
}
