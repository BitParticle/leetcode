package main
//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var res int
func maxDepth(root *TreeNode) int {
	if root==nil{
		return 0
	}
	leftMax := maxDepth(root.Left)
	rightMax := maxDepth(root.Right)
	res = max(leftMax, rightMax)+1
	return res
}

func max(a int, b int) int{
	if a>b{
		return a
	} else {
		return b
	}
}
//leetcode submit region end(Prohibit modification and deletion)
