package main

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for slow != nil {
		fast = fast.Next
		if fast == nil {
			return nil
		}
		fast = fast.Next
		if fast == nil {
			return nil
		}
		slow = slow.Next
		if slow == fast {
			fast = head
			for fast!=slow{
				fast=fast.Next
				slow=slow.Next
			}
			break
		}
	}
	return slow
}

//leetcode submit region end(Prohibit modification and deletion)
