package main

import (
	"leet"
)

func main() {
	//a1 :=[]int{6,4,5,0,4,4,9,4,1}
	//a2 := []int{3,8,8,0,3,0,1,4,8}
	//l1 := leet.MakeChain(a1)
	//l2 := leet.MakeChain(a2)
	//leet.AddTwoNumbers(l1,l2)
	s := "1231234545345"
	//for _,v := range s {
	//	println(fmt.Sprintf("%c",v))
	//}
	leet.LengthOfLongestSubstring(s)
}



func twoSum(nums []int, target int) []int {
	nums = quickSort(nums)
	for index, v := range nums {
		other := target - v
		v, i := leet.BinarySearch(nums, other, len(nums)-1, index+1)
		if v != -1 {
			return []int{index, i}
		}
	}
	return nil
}

func quickSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	left := 0
	mid := 0
	right := len(nums) - 1
	for {
		if nums[mid] < nums[left] {
			nums[mid], nums[left] = nums[left], nums[mid]
			mid=left
			left++
		}
		if nums[mid] > nums[right]{
			nums[mid], nums[right] = nums[right], nums[mid]
			mid=right
		}
		right--
		if left == right{
			break
		}
	}
	if len(nums) == 2 {
		return nums
	}
	quickSort(nums[:mid])
	quickSort(nums[mid:])
	return nums
}
