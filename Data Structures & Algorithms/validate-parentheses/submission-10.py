class Solution:
    def isValid(self, s: str) -> bool:
        if len(s) % 2 != 0:
            return False
        
        stack = []
        for b in s:
            if b == '(':
                stack.append(')')
            elif b == '[':
                stack.append(']')
            elif b == '{':
                stack.append('}')
            elif not bool(stack):
                return False
            elif b == ')':
                if b != stack.pop():
                    return False
            elif b == ']':
                if b != stack.pop():
                    return False
            elif b == '}':
                if b != stack.pop():
                    return False

        return not bool(stack)
            

