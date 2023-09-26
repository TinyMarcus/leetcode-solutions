
func twoSum(nums []int, target int) []int {
    numsMap := make(map[int]int, len(nums))
    var indexes []int

    for idx, val := range nums {
        if _, ok := numsMap[target - val]; ok {
            indexes = append(indexes, idx, numsMap[target - val])
        }

        numsMap[val] = idx
    }

    return indexes
}