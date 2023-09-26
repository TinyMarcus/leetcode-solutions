# Two Sum problem (1): https://leetcode.com/problems/two-sum/

class Solution:
    def twoSum(self, nums: list[int], target: int) -> list[int]:
        numsHash = dict()
        result = []

        for i in range(len(nums)):
            if nums[i] - target in numsHash:
                result.append(i)
                result.append(nums.index(target - nums[i]))

            numsHash[nums[i]] = i
        
        return result
