func findMaxAverage(nums []int, k int) float64 {
    start := 0
    maxAverage := float64(math.Inf(-1))
    sum := 0.0

    for end := 0; end < len(nums); end++ {
        sum += float64(nums[end])

        if end - start + 1 == k {
            maxAverage = getMax(maxAverage, float64(sum / float64(end - start + 1)))
            sum -= float64(nums[start])
            start++
        }
    }

    return maxAverage
}

func getMax(a, b float64) float64 {
    if a > b {
        return a
    } else {
        return b
    }
}