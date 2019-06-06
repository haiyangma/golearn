package main

import "fmt"

func pointerTest(pos *int ) int{
	*pos = 3
	return *pos
}

func swap(a,b *int) (int,int){
	*a,*b = *b,*a
	return *a,*b
}

func main() {
	var a int = 2
	fmt.Println(pointerTest(&a))
	fmt.Println(a)
	a,b := 1,2
	swap(&a,&b)
	fmt.Println(a,b)
}
