func findMaxConsecutiveOnes(nums []int) int {
	var out = 0
	var counter = 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			counter += 1
		}

		if nums[i] == 0 {
			if counter > out {
				out = counter
			}
			counter = 0
		}
	}

	if counter > out {
		return counter
	}

	return out
}
