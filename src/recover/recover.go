package main

import (
	"fmt"
)

func tryRecover(){
	defer func() {
		r := recover()
		if err,ok := r.(error);ok{
			fmt.Println("err occurd !",err)
		}else{
			panic(r)
		}
	}()
	//panic(errors.New("this is an error!"))
	//a := 5;
	//var b int = 0;
	//fmt.Println(a/b)
	panic(123)
}

func main() {
	tryRecover()
}
