package leet

import (
	"fmt"
	"strings"
)

func LengthOfLongestSubstring(s string) int {
	m := make(map[int]string)
	arr := toSArr(s)
	i := 0
	j := 0
	max := 0
	for {
		if i == len(s) {
			break
		}
		if v,ok := m[j];ok {
			if !strings.Contains(v,arr[i]) {
				m[j] = v+arr[i]
			}else {
				if j == (len(s) - 1) {
					break
				}
				j++
				i=j
				continue
			}
		} else {
			m[j] = arr[i]
		}
		i++
	}
	for _,v :=range m {
		if max < len(v) {
			max = len(v)
		}
	}
	return max
}

func toSArr(s string) []string {
	ret := make([]string,0)
	for _,c := range s {
		ret = append(ret,fmt.Sprintf("%c",c))
	}
	return ret
}
