func replaceElements(nums []int) []int {
	for i := 0; i < len(nums)-1; i++ {
		rightMax := nums[i+1]
		for j := i + 1; j < len(nums); j++ {
			if nums[j] > rightMax {
				rightMax = nums[j]
			}
		}
		nums[i] = rightMax
	}
	nums[len(nums)-1] = -1
	return nums
}
