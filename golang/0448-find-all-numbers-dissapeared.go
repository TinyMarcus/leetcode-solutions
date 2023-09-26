func findDisappearedNumbers(nums []int) []int {
	i := 0
	n := len(nums)
  
	for i < n {
		pos := nums[i]
  
		if nums[i] != nums[pos - 1] {
			nums[i], nums[pos - 1] = nums[pos - 1], nums[i]
	  	} else {
			i += 1
	  	}
	}
  
	var missed []int
  
	for i := range nums {
	  	if nums[i] != i + 1 {
			missed = append(missed, i + 1)
	  	}
	}
  
	return missed
}