func maxSubArray(nums []int) int {
    curSum := 0
    maxSum := nums[0]

    for _, val := range nums {
        curSum = max(curSum + val, val)
        maxSum = max(maxSum, curSum)
    }
    
    return maxSum
}

func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}