package easy

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
Merge two sorted linked lists and return it as a new list. The new list should
be made by splicing together the nodes of the first two lists.

Example:
	Input: 1->2->4, 1->3->4
	Output: 1->1->2->3->4->4
*/

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var res ListNode

	if l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			res.Val = l1.Val
			res.Next = mergeTwoLists(l1.Next, l2)
			return &res
		}
		res.Val = l2.Val
		res.Next = mergeTwoLists(l1, l2.Next)
		return &res

	}
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	return nil
}
