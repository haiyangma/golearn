package leet

import "fmt"

func Convert(s string, numRows int) string {
	return convert(s,numRows)
}

func convert(s string, numRows int) string {
	ret := make([]string,numRows)
	sl := len(s)
	y := numRows - 2
	if numRows < 3 {
		y = 0
	}
	n := sl/(numRows+y)
	if n > 0 {
		for i:=1;i<=n;i++ {
			target := i - 1
			start := target*numRows
			if target%2 == 0 {
				for ii:=0;ii<numRows;ii++ {
					ret[ii] = ret[ii]+fmt.Sprintf("%c",s[start])
					start++
				}
				if numRows>2 {
					for ii:=numRows-2;ii>=1;ii-- {
						ret[ii]=ret[ii]+fmt.Sprintf("%c",s[start])
						start++
					}
				}
			} else {
				for ii:=numRows-1;ii>=0;ii-- {
					ret[ii] = ret[ii]+fmt.Sprintf("%c",s[start])
					start++
				}
				if numRows>2 {
					for ii:=1;ii<=numRows-2;ii++ {
						ret[ii]=ret[ii]+fmt.Sprintf("%c",s[start])
						start++
					}
				}
			}
		}
	}
	start := sl-n
	first := 0
	last := numRows - 2
	if n % 2 == 0 {
		first = numRows - 2
		last = 0
	}
	if first < last {
		for i:=first;i<=last;i++ {
			ret[i]=ret[i]+fmt.Sprintf("%c",s[start])
			start++
		}
	} else {
		for i:=first;i<=last;i-- {
			ret[i]=ret[i]+fmt.Sprintf("%c",s[start])
			start++
		}
	}

	total := ""
	for _,ss := range ret {
		total = total+ss
	}
	return total
}