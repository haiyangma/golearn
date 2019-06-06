package main

import (
	"io/ioutil"
	"fmt"
)

func main() {
	const fileName = "abc.txt"
	contents,err := ioutil.ReadFile(fileName)
	if err != nil{
		fmt.Println(err)
	} else {
		fmt.Println(contents)
	}
}
