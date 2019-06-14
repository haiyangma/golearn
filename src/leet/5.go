package leet

import (
	"fmt"
	"strconv"
	"strings"
)

func LongestPalindrome(s string) string {
	return longestPalindrome(s)
}

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	dup := make(map[int32]struct{})
	set := make(map[int32][]int)
	for i, c := range s {
		if v, ok := set[c]; ok {
			set[c] = append(v, i)
		} else {
			set[c] = []int{i}
		}
	}
	ret := make(map[string]int)
	//start:=0
	//end:=0
	rets := ""
	max := 0
	maxleft := 0
	maxright := 0
	for i := len(s) - 1; i >= 0; i-- {
		if v, ok := set[int32(s[i])]; ok {
			if len(v) == 1 {
				continue
			}
			if _, ok := dup[int32(s[i])]; ok {
				continue
			}
		sameloop:
			for ii := 0; ii < len(v); ii++ {
				for jj := len(v) - 1; jj > ii; jj-- {
					left := v[ii]
					right := v[jj]
					if left > maxleft && right < maxright {
						continue
					}
					mid := (left + right) >> 1
					if (left+right)%2 == 0 {
						if len(s[left:]) > 2 && s[mid-1] != s[mid+1] {
							continue
						}
					} else {
						if len(s[left:]) > 1 && s[mid] != s[mid+1] {
							continue
						}
					}
					if right-left <= max {
						continue
					}
					ret[strconv.Itoa(int(v[ii]))+"-"+strconv.Itoa(int(v[jj]))] = 0
					for {
						if left != right {
							if s[left] == s[right] {
								ret[strconv.Itoa(int(v[ii]))+"-"+strconv.Itoa(int(v[jj]))] = v[jj] - v[ii]
								if left+1 == right || left+2 == right {
									if ret[strconv.Itoa(int(v[ii]))+"-"+strconv.Itoa(int(v[jj]))] > max {
										rets = strconv.Itoa(int(v[ii])) + "-" + strconv.Itoa(int(v[jj]))
										max = ret[strconv.Itoa(int(v[ii]))+"-"+strconv.Itoa(int(v[jj]))]
										maxleft = v[ii]
										maxright = v[jj]
										arr := strings.Split(rets, "-")
										left, _ := strconv.Atoi(arr[0])
										right, _ := strconv.Atoi(arr[1])
										println(s[left : right+1])
									}
									break sameloop
								}
								left++
								right--
							} else {
								break
							}
						} else {
							if ret[strconv.Itoa(int(v[ii]))+"-"+strconv.Itoa(int(v[jj]))] > max {
								rets = strconv.Itoa(int(v[ii])) + "-" + strconv.Itoa(int(v[jj]))
								max = ret[strconv.Itoa(int(v[ii]))+"-"+strconv.Itoa(int(v[jj]))]
								maxleft = v[ii]
								maxright = v[jj]
							}
							break
						}
					}
				}
				dup[int32(s[i])] = struct{}{}
			}

		}
	}
	if rets == "" {
		return fmt.Sprintf("%c", s[0])
	}
	arr := strings.Split(rets, "-")
	left, _ := strconv.Atoi(arr[0])
	right, _ := strconv.Atoi(arr[1])
	return s[left : right+1]
}

