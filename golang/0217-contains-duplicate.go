func containsDuplicate(nums []int) bool {
    numsMap := map[int]int{}

    for _, val := range nums {
        if _, ok := numsMap[val]; ok == true {
            return true
        }

        numsMap[val] = 1
    }

    return false
}