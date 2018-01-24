package easy

import (
	"reflect"
	"testing"
)

// make a linked ListNode from the array
func makeListNode(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	var (
		head = new(ListNode)
		tmp  = head
	)
	for i, val := range vals {
		tmp.Val = val
		if i == len(vals)-1 {
			tmp.Next = nil
		} else {
			tmp.Next = new(ListNode)
			tmp = tmp.Next
		}
	}
	return head
}

// make a slice from the linked ListNode
func (ln *ListNode) toList() []int {
	var (
		tmp = ln
		res []int
	)
	for tmp != nil {
		res = append(res, tmp.Val)
		tmp = tmp.Next
	}
	return res
}

func TestMergeTwoLists(t *testing.T) {
	cases := []struct {
		l1  []int
		l2  []int
		out []int
	}{
		{
			[]int{0, 1, 2},
			[]int{0},
			[]int{0, 0, 1, 2},
		},
		{
			[]int{0, 0, 1, 1},
			[]int{1, 2, 3},
			[]int{0, 0, 1, 1, 1, 2, 3},
		},
		{
			[]int{},
			[]int{},
			[]int{},
		},
		{
			[]int{0},
			[]int{},
			[]int{0},
		},
		{
			[]int{99},
			[]int{1},
			[]int{1, 99},
		},
		{
			[]int{},
			[]int{0},
			[]int{0},
		},
	}
	for _, c := range cases {
		var l1, l2 = makeListNode(c.l1), makeListNode(c.l2)
		var out = mergeTwoLists(l1, l2).toList()
		if len(out) != 0 && len(c.out) != 0 && !reflect.DeepEqual(out, c.out) {
			t.Errorf("mergeTwoList(%v,%v)=%v, expected %v", c.l1, c.l2, out, c.out)
		}
	}
}
