class Solution:
    def findMaxConsecutiveOnes(self, nums: List[int]) -> int:
        best = 0
        curr = 0
        for val in nums:
            if val == 1:
                curr += 1
            else:
                best = max(best, curr)
                curr = 0
        return max(best, curr)

