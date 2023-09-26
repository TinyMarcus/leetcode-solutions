func getMin(a, b int) int {
    if a > b {
        return b
    } else {
        return a
    }
}

func getMax(a, b int) int {
    if a < b {
        return b
    } else {
        return a
    }
}

func maxArea(height []int) int {
    start := 0
    end := len(height) - 1
    maxArea := 0

    for start < end {
        maxArea = getMax(maxArea, getMin(height[start], height[end]) * (end - start))

        if height[start] < height[end] {
            start++
        } else {
            end--
        }
    }

    return maxArea
}