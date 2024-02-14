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
func levelTraverse(root *TreeNode, level int, res *[][]int) {
	if root == nil {
		return
	}
	if len(*res) < level {
		*res = append(*res, nil)
	}
	(*res)[level-1] = append((*res)[level-1], root.Val)
	levelTraverse(root.Left, level+1, res)
	levelTraverse(root.Right, level+1, res)
}
func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	levelTraverse(root, 1, &res)
	return res
}

//leetcode submit region end(Prohibit modification and deletion)
