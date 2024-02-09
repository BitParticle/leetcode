package main

//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for slow != nil {
		fast = fast.Next
		if fast == nil{
			return false
		}
		fast = fast.Next
		if fast==nil {
			return false
		}
		slow = slow.Next
		if slow == fast{
			return true
		}
	}
	return false
}

//leetcode submit region end(Prohibit modification and deletion)
