package main

type ListNode struct {
	   Val int
	   Next *ListNode
 }
//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
   dummy1 := headA
   dummy2 := headB
   countA,countB := 0,0
   for dummy1!=nil {
	   countA++
	   dummy1 = dummy1.Next
   }
   for dummy2!=nil {
	   countB++
	   dummy2 = dummy2.Next
   }
   if countA>countB {
	  step := countA - countB
	  for i:=0;i<step;i++ {
		  headA=headA.Next
	  }
   }else{
	   step := countB-countA
	   for i:=0;i<step;i++{
		   headB=headB.Next
	   }
   }
   for headA!=headB{
	   headA=headA.Next
	   headB=headB.Next
   }
   return headA
}
//leetcode submit region end(Prohibit modification and deletion)
