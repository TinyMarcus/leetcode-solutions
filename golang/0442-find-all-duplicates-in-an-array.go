func findDuplicates(nums []int) []int {
    numsMap := map[int]int{}

	for _, val := range nums {
		if _, ok := numsMap[val]; ok == false {
			numsMap[val] = 1
		} else {
			numsMap[val] += 1
		}
	}

	var res []int

	for key, val := range numsMap {
		if val == 2 {
			res = append(res, key)
		}
	}

	return res
}