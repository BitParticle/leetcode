//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	p1 := &ListNode{-1, nil}
	p2 := &ListNode{-1, nil}
	p1Dummy := p1
	p2Dummy := p2
	p := head
	for p!=nil {
		if p.Val<x {
			p1.Next = p
			p1 = p1.Next
		}else {
			p2.Next = p
			p2 = p2.Next
		}
		temp := p.Next
		p.Next = nil
		p = temp
	}
	p1.Next = p2Dummy.Next
	return p1Dummy.Next
}
//leetcode submit region end(Prohibit modification and deletion)
