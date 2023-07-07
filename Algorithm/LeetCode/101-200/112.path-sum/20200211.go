/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var hasSum bool

func addPathSum(t *TreeNode, sum, sumNow int)  {
    if t == nil {
        return
    } 
    
    val := t.Val
    sumNow += val
    if sum == sumNow && isLeaf(t){
            hasSum = true
    } else {
        addPathSum(t.Left, sum, sumNow)
        addPathSum(t.Right, sum, sumNow)
    }
}

func isLeaf(t *TreeNode) bool {
    if t != nil && t.Left ==nil && t.Right == nil {
        return true
    }   
    return false
}

func hasPathSum(root *TreeNode, sum int) bool {
    hasSum = false
    if root == nil {
        return false
    }

    sumNow := root.Val
    if isLeaf(root) && sum == sumNow {
        hasSum = true
    }
    
    addPathSum(root.Left, sum, sumNow)
    addPathSum(root.Right, sum, sumNow)

        return hasSum
}
