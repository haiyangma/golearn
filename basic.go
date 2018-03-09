package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	aa = 3
	bb = 4
	cc = 5
)

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func initialValue() {
	var a int = 3
	var s string = "abc"
	fmt.Println(a, s)
}

func variableTypeDeduction() {
	var a, b, c, d = 3, 4, 5, 6
	fmt.Println(a, b, c, d)
}

func variableShorter() {
	a, b, c, d := 3, 4, "def", false
	b = 5
	fmt.Println(a, b, c, d)
}

func complxTest() {
	fmt.Printf("%.3f", cmplx.Pow(math.E, 1i*math.Pi))
}

func triangle() {
	var a, c int = 3, 4
	var d int
	d = int(math.Sqrt(float64(a*a + c*c)))
	fmt.Println(d)
}

func consts() {
	const fileName string = "abc.txt"
	const a, b = 3, 4
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(fileName, c)
}

func enumsTest() {
	const (
		cpp = iota
		_
		java
		python
		golang
	)
	fmt.Println(cpp, java, python, golang)
}

func main() {
	//fmt.Println("Hello world !")
	//variableZeroValue()
	//initialValue()
	//variableTypeDeduction()
	//variableShorter()
	//complxTest()
	//fmt.Println("\n")
	//triangle()
	//consts()
	//enumsTest()
}
