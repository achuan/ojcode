package main

//https://leetcode-cn.com/problems/recover-binary-search-tree/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func recoverTree(root *TreeNode) {
	var prev, t1, t2 *TreeNode
	var inorder func (root * TreeNode)
	inorder = func (root *TreeNode){
		if root == nil {
			return
		}
		inorder(root.Left)
		if prev != nil && prev.Val > root.Val{
			if t1 == nil {
				t1 = prev
			}
			t2 = root
		}
		prev = root
		inorder(root.Right)
	}
	inorder(root)
	t1.Val, t2.Val = t2.Val, t1.Val
	return
}

func createTree(root **TreeNode, i int, array []int) {
	size := len(array)
	if i >= size || array[i] == -1 {
		return
	}
	if *root == nil {
		*root = new(TreeNode)
	}

	(*root).Val = array[i]
	createTree(&((*root).Left), 2*i+1, array)
	createTree(&((*root).Right), 2*i+2, array)
}

func initTree(array []int) *TreeNode {
	root := &TreeNode{0, nil, nil}
	createTree(&root, 0, array)
	return root
}

func main(){
	testArray := [][]int{
		{1, 3, -1, -1, 2},
		{3, 1, 4, -1, -1, 2, -1},
		{4, 1, 5, 3, -1, -1, -1, -1, 2},
	}
	for _, array := range testArray {
		tree := initTree(array)
		recoverTree(tree)
	}
}
