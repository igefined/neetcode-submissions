class Solution:
    def replaceElements(self, arr: List[int]) -> List[int]:
        prev = -1
        best = prev

        i = len(arr) - 1
        while i >= 0:
            best = max(best, prev)
            prev = arr[i]
            arr[i] = best
            i -= 1
        return arr

