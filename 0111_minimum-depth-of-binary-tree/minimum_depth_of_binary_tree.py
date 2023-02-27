# Minimum depth of binary tree (111): https://leetcode.com/problems/minimum-depth-of-binary-tree/
# Pattern: BFS

from collections import deque

# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right

class Solution:
    def minDepth(self, root: TreeNode) -> int:
        if not root:
            return 0

        q = deque()
        q.append(root)
        levelCount = 0

        while len(q):
            size = len(q)
            levelCount += 1

            for _ in range(size):
                node = q.popleft()

                if not node.left and not node.right:
                    return levelCount
                
                if node.left:
                    q.append(node.left)

                if node.right:
                    q.append(node.right)
