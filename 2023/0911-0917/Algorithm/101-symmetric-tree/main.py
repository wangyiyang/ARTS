# Definition for a binary tree node.
class TreeNode(object):
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

class Solution(object):
    def isSymmetric(self, root):
        """
        :type root: TreeNode
        :rtype: bool
        """
        return self.checkSymmetric(root, root)

    def checkSymmetric(self, left, right):
        if not left and not right:
            return True
        elif not left or not right:
            return False
        return left.val == right.val and self.checkSymmetric(left.left, right.right) and self.checkSymmetric(left.right, right.left)


if __name__ == '__main__':
    root = TreeNode(1)
    root.left = TreeNode(2)
    root.left.left = TreeNode(3)
    root.right = TreeNode(2)
    root.right.right = TreeNode(3)
    s = Solution()
    print(s.isSymmetric(root))
        
