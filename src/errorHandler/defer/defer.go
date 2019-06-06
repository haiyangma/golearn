package main

import (
	"bufio"
	"fmt"
	"os"
	"functional/fib"
)

func tryDefer(){
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
}

func writeFile(fineName string){
		file,err := os.OpenFile(fineName, os.O_EXCL|os.O_CREATE, 0666);
		if err!=nil{
			if pathError,ok := err.(*os.PathError);!ok{
				panic(err)
			}else{
				fmt.Print(pathError)
			}
			return
		}
		defer file.Close()

		f:=fib.Fibonacci1()
		writer := bufio.NewWriter(file)
		defer writer.Flush()
		for i:=0;i<20;i++{
			fmt.Fprintln(writer,f())
		}
}

func main() {
	tryDefer()
	writeFile("fib.txt")
}
