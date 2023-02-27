# Climbing stairs (70): https://leetcode.com/problems/climbing-stairs
# Used fibonacci

class Solution:
    def climbStairs(self, n: int) -> int:
        first, second = 1, 1

        for i in range(n - 1):
            temp = first
            first = first + second
            second = temp
        
        return first