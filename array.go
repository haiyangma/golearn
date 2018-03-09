package main

import "fmt"

func printArr(arr *[3]int){
	for i,v := range *arr{
		fmt.Println(i,v)
	}
}

func ipdateSlice(arr []int){
	arr[0] = 100
}

func main() {
	var arr1 [5]int
	arr2 := [3]int{1,2,3}
	arr3 := [...]int{1}

	var grid [4][5]bool

	fmt.Println(arr1,arr2,arr3,grid)

	for i := 0;i<len(arr3);i++{
		fmt.Println(arr3[i])
	}

	for _,v := range arr3{
		fmt.Println(v)
	}

	printArr(&arr2)

	arr := [...]int{1,3,4,5,6,78,9}
	s := arr[2:6]
	fmt.Println("arr[2:6] = ",s)
	fmt.Println("arr[:6] = ",arr[:6])
	fmt.Println("arr[2:] = ",arr[2:])
	fmt.Println("arr[:] = ",arr[:])

	ipdateSlice(arr[2:]);
	fmt.Println(arr)


	arr4 := [...]int{0,1,2,3,4,5,6,7}
	s1 := arr4[2:6]
	s2 := s1[3:5]
	fmt.Println(s1)
	fmt.Println(s2,len(s2),cap(s2))
	fmt.Println(s2[1:3])

	s3 := append(s2,10)
	s4 := append(s3,11)
	s5 := append(s4,12)
	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Println(s5)
	fmt.Println(arr4)
}
