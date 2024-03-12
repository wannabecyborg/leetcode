package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func InsertionSortListFirstSolution(head *ListNode) *ListNode {
	for n := head.Next; n != nil; n = n.Next {
		if n.Val < head.Val {
			currVal := head.Val
			head.Val = n.Val
			for next := head.Next; next != n.Next; next = next.Next {
				c := next.Val
				next.Val = currVal
				currVal = c
			}
		} else {
			for j := head.Next; j != n; j = j.Next {
				if n.Val < j.Val {
					c := j.Val
					j.Val = n.Val
					n.Val = c
				}
			}
		}
	}
	return head
}

func InsertionSortListBestSolution(head *ListNode) *ListNode {
	var newhead *ListNode
	curr := head
	for curr != nil {
		temp := curr.Next
		newhead = sortedInsert(curr, newhead)
		curr = temp
	}
	return newhead
}

func sortedInsert(curr *ListNode, newHead *ListNode) *ListNode {
	if newHead == nil || newHead.Val >= curr.Val {
		curr.Next = newHead
		newHead = curr
	} else {
		temp := newHead
		for temp.Next != nil && temp.Val < curr.Val {
			temp = temp.Next
		}
		curr.Next = temp.Next
		temp.Next = curr
	}
	return newHead
}

func InsertionSortSimpleArray(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		for j := i; j != 0 && arr[j-1] > arr[j]; j-- {
			x := arr[j]
			arr[j] = arr[j-1]
			arr[j-1] = x
		}
	}
	return arr
}
