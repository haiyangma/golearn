package main

import "fmt"


func lengthOfNonRepetingSubStr(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxLength := 0
	for i,ch := range []rune(s) {

		if lastI,ok := lastOccurred[ch];ok && lastI >= start {
			start = lastOccurred[ch] + 1
		}
		if i - start + 1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}

func main() {

	m := map[string]string {
		"name":"mhy",
		"course":"golang",
		"quality":"notbad",
	}

	m2 := make(map[string]int)

	var m3 map[string]int

	fmt.Println(m,m2,m3)

	fmt.Println("traversing map")
	for k,v := range m {
		fmt.Println(k,v)
	}

	fmt.Println("getting values")
	name,ok := m["name"]
	fmt.Println(name,ok)

	delete(m,"name")
	name,ok = m["name"]
	fmt.Println(name,ok)


	fmt.Println(lengthOfNonRepetingSubStr("我爱我你我中我爱我你我国你"))

}
