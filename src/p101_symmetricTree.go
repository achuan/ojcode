package main

//https://leetcode-cn.com/problems/symmetric-tree/

func helper(tree1, tree2 *TreeNode)bool{
	if tree1 == nil && tree2 == nil {
		return true
	}
	if tree1 == nil || tree2 == nil || tree1.Val != tree2.Val {
		return false
	}
	return helper(tree1.Left, tree2.Right) && helper(tree1.Right, tree2.Left)
}

func isSymmetric(root *TreeNode) bool {
	if root == nil{
		return true
	}
	return helper(root.Left, root.Right)
}

func main(){

}
