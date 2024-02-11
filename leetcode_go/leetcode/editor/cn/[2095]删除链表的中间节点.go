package main

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func deleteMiddle(head *ListNode) *ListNode {
	slow, fast := head, head
	len := 0
	for fast != nil {
		len++
		fast = fast.Next
	}
	if len == 1 {
		return nil
	} else {
		if len%2==0{
			len=len+1
		}
		for i := 1; i < len/2; i++ {
			slow = slow.Next
		}
		temp := slow.Next
		slow.Next = temp.Next
		return head
	}
}

//leetcode submit region end(Prohibit modification and deletion)
