package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"math/cmplx"
	"os"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"unicode/utf8"
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

	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)

	fmt.Println(cpp, java, python, golang)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func grade(score int) string {
	g := ""

	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong score %d", score))
	case score < 60:
		g = "F"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score <= 100:
		g = "A"
	}

	return g
}

func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile(filename string) {
	file, err := os.Open(filename)

	if err != nil {
		panic(err)
	}

	printFileContents(file)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func eval(a, b int, op string) (result int, err error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		return 0, errors.New("unsupported operation: " + op)
	}
}

func div(a, b int) (q, r int) {
	return a / b, a % b
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args (%d, %d)\n", opName, a, b)

	return op(a, b)
}

func sum(numbers ...int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func swap(a, b *int) {
	*b, *a = *a, *b
}

func printArray(arr *[5]int) {
	arr[0] = 100
	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func updateSlice(s []int) {
	s[0] = 100
}

func printSlice(s []int) {
	fmt.Printf("v=%v, len=%d, cap=%d\n", s, len(s), cap(s))
}

func slice() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s := arr[2:6]
	fmt.Println(s)
	updateSlice(s)
	fmt.Println(s)
	fmt.Println(arr)

	s2 := s[3:5]
	fmt.Println(s2)

	fmt.Println(len(s), cap(s))

	s3 := append(s2, 10)
	s4 := append(s3, 11)
	s5 := append(s4, 12)
	s6 := append(s5, 13)

	fmt.Println(s3, s4, s5, s6)
	fmt.Println(arr)

	var sl []int

	for i := 0; i < 100; i++ {
		printSlice(sl)
		sl = append(sl, 2*i+1)
	}

	s7 := []int{2, 4, 6, 8}
	s8 := make([]int, 16)
	s9 := make([]int, 10, 32)

	printSlice(s7)
	printSlice(s8)
	printSlice(s9)

	copy(s9, s7)
	printSlice(s9)

	s9 = append(s9[:3], s9[4:]...)
	printSlice(s9)

	front := s9[0]
	s9 = s9[1:]

	tail := s9[len(s9)-1]
	s9 = s9[:len(s9)-1]

	fmt.Println(front, tail)
	printSlice(s9)
}

func maps() {
	m := map[string]string{
		"name":    "liwenxiang",
		"course":  "golang",
		"site":    "imooc",
		"quality": "notbad",
	}

	m2 := make(map[string]int)

	fmt.Println(m)
	fmt.Println(m2)

	for k, v := range m {
		fmt.Println(k, v)
	}

	if name, ok := m["course1"]; ok {
		fmt.Println(name)
	} else {
		fmt.Println("key does not exist")
	}
}

func nonRepeating() {
	lengthOfNonRepeatingSubStr("abcddbcaefgh")
	lengthOfNonRepeatingSubStr("我爱慕课网!!")
}

func lengthOfNonRepeatingSubStr(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0

	for i, ch := range []rune(s) {
		if lastI, ok := lastOccurred[ch]; ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}

	fmt.Println("Max length:", maxLength)

	return maxLength
}

func myStrings() {
	s := "Yes我爱慕课网!"
	fmt.Println(len(s))
	for _, b := range []byte(s) {
		fmt.Printf("%X ", b)
	}
	fmt.Println()
	for i, ch := range s {
		fmt.Printf("(%d %X)", i, ch)
	}
	fmt.Println()
	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	bytes := []byte(s)
	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}
	fmt.Println()

	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c) ", i, ch)
	}
}

func main() {
	fmt.Println("Hello World!")
	variableZeroValue()
	euler()
	enums()
	fmt.Println(
		grade(0),
		grade(100),
		grade(60),
		grade(59),
	)
	fmt.Println(convertToBin(14))
	printFile("abc.txt")

	for i := 0; i < 10; i++ {
		println(i)
	}

	if a, err := eval(13, 3, "x"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(a)
	}

	fmt.Println(div(13, 3))

	fmt.Println(apply(func(a, b int) int {
		return a - b
	}, 13, 3))

	fmt.Println(sum(1, 2, 3, 4, 5))

	a, b := 3, 4
	swap(&a, &b)
	fmt.Println(a, b)

	var arr1 [5]int
	arr2 := [3]int{2, 3, 4}
	arr3 := [...]int{2, 3, 4, 5, 6}
	var arr4 [4][5]bool

	fmt.Println(arr1, arr2, arr3, arr4)

	printArray(&arr1)
	printArray(&arr3)
	fmt.Println(arr1, arr3)

	slice()

	maps()

	nonRepeating()

	myStrings()

	fmt.Println()

	s := `aaa "bbb"
	kkk
	ddd
	1231222`

	printFileContents(strings.NewReader(s))
}
