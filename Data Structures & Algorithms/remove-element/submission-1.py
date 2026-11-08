class Solution:
    def removeElement(self, nums: List[int], val: int) -> int:
        count = len(nums)
        prev = 0
        for i, value in enumerate(nums):
            nums[i - prev] = nums[i]
            if val == value:
                count -= 1
                prev += 1
                continue
        return count
