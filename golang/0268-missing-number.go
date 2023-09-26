func missingNumber(nums []int) int {
    factSum := findSum(nums)
    theoreticalSum := 0
    n := len(nums)

    for i := 0; i <= n; i++ {
        theoreticalSum += i
    }

    return theoreticalSum - factSum
}

func findSum(nums []int) int {
    sum := 0

    for _, val := range nums {
        sum += val
    }

    return sum
}