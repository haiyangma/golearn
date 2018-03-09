package main

import (
	"fmt"
	"reflect"
	"runtime"
	"math"
)

func eval(a,b int,op string) int {
	switch op{
	case "+":
		return a+b
	case "-":
		return a-b
	case "*":
		return a*b
	case "/":
		return a/b
	default:
		panic("unsupported opration: "+op)
	}
}

func div(a,b int)(q,r int){
	q = a/b
	r = a%b
	return
}

func apply(op func(int,int) int ,a,b int) int{
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Printf("calling function %s with args (%d %d) \n ",opName,a,b)
	return op(a,b)
}

func pow(a,b int) int {
	return int(math.Pow(float64(a),float64(b)))
}

func sum(numbers ... int) int {
	s := 0
	for i := range numbers {
		s += numbers[i]
	}
	return s
}

func main() {
	fmt.Println(eval(3,4,"+"))
	r,q := div(13,5)
	fmt.Println(r,q)
	//fmt.Println(apply(pow,3,4))
	fmt.Println(apply(
		func(a int,b int) int{
		//p := reflect.ValueOf(op).Pointer()
		//opName := runtime.FuncForPC(p).Name()
		//fmt.Printf("calling function %s with args (%d %d) \n ",runtime.FuncForPC(reflect.ValueOf().Pointer()).Name(),a,b)
		return int(math.Pow(float64(a),float64(b)))
	},3,4))
	fmt.Println(sum(1,2,3,4))

}
