package leet


type ListNode struct {
    Val int
    Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	add := 0
	return sum(l1,l2,&add)
}

func sum(l1 *ListNode, l2 *ListNode,add *int) *ListNode {
	s := 0
	if l1 != nil {
		s+=l1.Val
	}
	if l2 != nil {
		s+=l2.Val
	}
	s+=*add
	toAdd := 0
	v := s
	if s>9 {
		toAdd = 1
		v = s - 10
	}
	result := ListNode{v,nil}
	if l1 == nil && l2 == nil {
		return &result
	}
	var targetL1 *ListNode
	var targetL2 *ListNode
	if l1!=nil {
		targetL1 = l1.Next
	}
	if l2 != nil {
		targetL2 = l2.Next
	}
	if toAdd >0 || targetL2!=nil || targetL1 != nil{
		result.Next = sum(targetL1,targetL2,&toAdd)
	}
	return &result
}

func MakeChain(nums []int) *ListNode {

	ret := ListNode{}
	ret.Val = nums[0]
	if len(nums) == 1 {
		return &ret
	}
	ret.Next = MakeChain(nums[1:])
	return &ret
}


