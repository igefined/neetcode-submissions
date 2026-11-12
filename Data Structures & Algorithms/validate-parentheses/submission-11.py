class Solution:
    def isValid(self, s: str) -> bool:
        stack = []
        chars = {'}':'{', ')':'(', ']':'['}

        for c in s:
            if c in chars:
                if stack and stack[-1] == chars[c]:
                    stack.pop()
                else:
                    return False
            else:
                stack.append(c)
        return True if not stack else False
            

