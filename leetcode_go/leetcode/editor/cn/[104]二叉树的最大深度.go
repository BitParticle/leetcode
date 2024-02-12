package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(root *TreeNode) int {
	res := 0
	depth := 0
	traverse(root, &res, &depth)
	return res
}

func traverse(root *TreeNode, res *int, depth *int) {
	if root == nil {
		return
	}
	*depth++
	if root.Right == nil && root.Left == nil {
		*res = max(*res, *depth)
	}
	traverse(root.Left, res, depth)
	traverse(root.Right, res, depth)
	*depth--
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
