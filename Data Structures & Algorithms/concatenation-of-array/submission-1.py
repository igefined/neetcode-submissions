class Solution:
    def getConcatenation(self, nums: List[int]) -> List[int]:
        n = len(nums)
        ans = [0] * n * 2
        i = 0
        while i < n:
            ans[i] = nums[i]
            ans[i + n] = nums[i]
            i += 1
        return ans
