class Solution:
    def replaceElements(self, arr: List[int]) -> List[int]:
        prev = -1
        best = prev

        for i in range(len(arr) - 1, -1, -1):
            best = max(best, prev)
            prev = arr[i]
            arr[i] = best
        return arr

