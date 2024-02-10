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
func midTraverse(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	midTraverse(root.Left, result)
	*result = append(*result, root.Val)
	midTraverse(root.Right, result)
}
func inorderTraversal(root *TreeNode) []int {
	var result []int
	midTraverse(root, &result)
	return result
}

//leetcode submit region end(Prohibit modification and deletion)
