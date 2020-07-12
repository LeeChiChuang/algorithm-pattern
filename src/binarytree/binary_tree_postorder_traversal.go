package binarytree

// 145. 二叉树的后序遍历 https://leetcode-cn.com/problems/binary-tree-postorder-traversal
// 给定一个二叉树，返回它的 后序 遍历。
//
// 示例:
//
// 输入: [1,null,2,3]
// 1
// \
// 2
// /
// 3
//
// 输出: [3,2,1]
// 进阶: 递归算法很简单，你可以通过迭代算法完成吗？

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func PostorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	postorder(root, &res)
	return res
}

func postorder(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}

	postorder(root.Left, res)
	postorder(root.Right, res)
	*res = append(*res, root.Val)
}

// 迭代 栈模仿递归
func PostorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	call := make([]*TreeNode, 0)
	res := make([]int, 0)
	var t *TreeNode
	call = append(call, root)
	for len(call) != 0 {
		t = call[len(call)-1]
		call = call[:len(call)-1]
		if t != nil {
			call = append(call, t)
			call = append(call, nil)
			if t.Right != nil {
				call = append(call, t.Right)
			}
			if t.Left != nil {
				call = append(call, t.Left)
			}
		} else {
			res = append(res, call[len(call)-1].Val)
			call = call[:len(call)-1]
		}
	}

	return res
}
