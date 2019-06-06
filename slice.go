package main

import "fmt"

func printSlice(s []int){
	fmt.Println(s)
	fmt.Printf("len = %d, cap = %d\n",len(s),cap(s))
}

func main() {

	var s []int // zero value for slice is nil
	fmt.Println(s == nil)

	for i := 0;i<100;i++{
		s = append(s,2*i +1)
		printSlice(s)
	}

	fmt.Println(s)

	s1 := []int{2,4,6,8}
	printSlice(s1)

	s2 := make([]int,16)

	s3 := make([]int,16,32)

	printSlice(s2)
	printSlice(s3)

	fmt.Println("copying slice")
	copy(s2,s1)
	printSlice(s2)

	fmt.Println("deleting elements from slice")
	s2 = append(s2[:3],s2[4:]...)
	printSlice(s2)

	fmt.Println("poping from front")
	front := s2[0]
	s2 = s2[1:]
	fmt.Println("fromt = ",front)
	printSlice(s2)
	fmt.Println("poping from end")
	end:=s2[len(s2)-1]
	fmt.Println("end = ",end)
	s2 = s2[:len(s2)-1]
	printSlice(s2)
}
