package main

import (
	"fmt"
	"math"
)

func triangle() {
	var a, b int = 3, 4
	c := calcTriangle(a, b)
	fmt.Println(c)
}

func calcTriangle(a, b int) int {
	return int(math.Sqrt(float64(a*a + b*b)))
}

func main() {
	triangle()
}
