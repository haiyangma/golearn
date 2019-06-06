package fib

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type Fibonacci func() (int,Fibonacci)


func fibonacci2(base1 int, base2 int) Fibonacci{
	return func() (int,Fibonacci){
		return base1+base2,fibonacci2(base2,base1+base2)
	}
}

func Fibonacci1() func() int{
	a,b := 0,1
	return func() int {
		a,b = b,a+b
		return a
	}
}

func (node *intGen) PrintFile(){
	scanner := bufio.NewScanner(node)
	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
}


type intGen func() int

func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	if next>=10000{
		return 0,io.EOF
	}
	s := fmt.Sprintf("%d\n",next)
	return strings.NewReader(s).Read(p)
}

func main() {

	f := fibonacci2(0,1)
	for i:=1;i<10;i++{
		var s int
		s,f = f()
		fmt.Println(s) //1
	}

	//f1 := fibonacci1()
	//f1.PrintFile()

	//s,f := fibonacci2(0,1)
	//fmt.Println(f()) //1
	//fmt.Println(f()) //1
	//fmt.Println(f()) //2
	//fmt.Println(f()) //3
	//fmt.Println(f()) //5
	//fmt.Println(f()) //8
	//fmt.Println(f()) //13
	//fmt.Println(f()) //21

}
