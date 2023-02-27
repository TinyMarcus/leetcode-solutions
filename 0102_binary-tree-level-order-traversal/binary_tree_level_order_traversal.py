# Binary tree level order traversal (102): https://leetcode.com/problems/binary-tree-level-order-traversal/
# Pattern: BFS

from collections import deque

# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

class Solution:
    def levelOrder(self, root: TreeNode) -> list[list[int]]:
        if not root:
            return []
        
        q = deque()
        q.append(root)
        big_list = []

        while len(q):
            size = len(q)
            small_list = []

            for _ in range(size):
                node = q.popleft()
                small_list.append(node.val)

                if node.left:
                    q.append(node.left)
                
                if node.right:
                    q.append(node.right)
                
            big_list.append(small_list)
        
        return big_list
                


        
