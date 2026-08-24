func replaceElements(arr []int) []int {
	maxRight := -1
	for i := len(arr) - 1; i >= 0; i-- {
		cur := arr[i]
		arr[i] = maxRight
		if cur > maxRight {
			maxRight = cur
		}
	}
	return arr
}
