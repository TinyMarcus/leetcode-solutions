# Container with most water (11): https://leetcode.com/problems/container-with-most-water/

class Solution:
    def maxArea(self, height: list[int]) -> int:
        firstPointer = 0
        secondPointer = len(height) - 1
        maxWater = 0

        while (firstPointer < secondPointer):
            maxWater = max(maxWater, (secondPointer - firstPointer) * min(height[firstPointer], height[secondPointer]))

            if (height[firstPointer] < height[secondPointer]):
                firstPointer += 1
            else:
                secondPointer -= 1

        return maxWater
