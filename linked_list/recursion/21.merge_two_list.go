package linkedlist

import linkedlist "go-programming-techniques/linked_list/math"

func MergeTwoLists(list1 *linkedlist.ListNode, list2 *linkedlist.ListNode) *linkedlist.ListNode {
	var result linkedlist.ListNode
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	if list1.Val < list2.Val {
		result.Val = list1.Val
		result.Next = MergeTwoLists(list1.Next, list2)
	} else {
		result.Val = list2.Val
		result.Next = MergeTwoLists(list1, list2.Next)
	}
	return &result
}
