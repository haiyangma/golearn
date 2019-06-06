package leet

import "math"

//func twoSum(nums []int, target int) []int {
//
//}
//
//func reduceNums(nums []int, target int) []int {
//	index := int(math.Round(float64(len(nums) / 2)))
//
//}


func twoSum1(nums []int, target int) []int {
	m := make(map[int]int)
	for index,v := range nums {
		other := target - v
		if t,ok := m[v];ok {
			return []int{t,index}
		}
		m[other]=index
	}
	return nil
}

func BinarySearch(nums []int, target int, high int, low int) (int,int) {
	mid := int(math.Ceil(float64((high + low) / 2)))
	if high - low == 1 && nums[low] != target{
		return -1,high
	}
	if nums[mid] > target {
		return BinarySearch(nums, target, low, mid-1)
	} else if nums[mid] < target {
		return BinarySearch(nums, target, mid+1, high)
	} else {
		return nums[mid],mid
	}
}
