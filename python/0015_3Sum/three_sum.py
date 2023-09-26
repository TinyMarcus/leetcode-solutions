# 3Sum (15): https://leetcode.com/problems/3sum/
# Pattern: two pointers

class Solution:
    def threeSum(self, nums: list[int]) -> list[list[int]]:
        nums = sorted(nums)
        result = set()
        n = len(nums)

        for i in range(n):
            target = -nums[i]
            left = i + 1
            right = n - 1

            while left < right:
                curSum = nums[left] + nums[right]

                if curSum == target:
                    result.add((nums[i], nums[left], nums[right]))

                if curSum < target:
                    left += 1
                else:
                    right -= 1

        return result