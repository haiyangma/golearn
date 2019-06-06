package main

import (
	"fmt"
	"strconv"
	"os"
	"bufio"
)

func convertToBin(n int) string {
	result := ""
	for ; n>0;n/=2{
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return  result
}

func printFile(fileName string){
	file,err := os.Open(fileName)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
}

func forTest(){
	for{
		fmt.Println("abc")
	}
}

func main() {
	fmt.Println(
		convertToBin(5),
		convertToBin(13),
	)

	printFile("abc.txt")
	forTest()
}
