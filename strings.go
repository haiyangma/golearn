package main

import (
	"fmt"
	"unicode/utf8"
	"strings"
)

func main() {

	s := "我爱编程"
	fmt.Println(len(s))

	for i,ch := range s {
		fmt.Printf("%d  %d \n",i,rune(ch))
	}

	bytes := []byte(s)
	for len(bytes)>0 {
		ch,size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c \n",ch)
	}

	var runeArr []rune = []rune(s);

	for i,ch := range runeArr {
		fmt.Printf("%d %c \n",i,ch)
	}

	fmt.Println(runeArr)
	fmt.Println(&s)
	str := "a 2 b 3 d 4"
	fmt.Println(strings.Fields(str))
}
