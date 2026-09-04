func hasDuplicate(nums []int) bool {
    set := make(map[int]struct{}, len(nums))
	for i := 0; i < len(nums); i++ {
		if _, ok := set[nums[i]]; ok {
			return true
		}
		set[nums[i]] = struct{}{}
	}
	return false
}
