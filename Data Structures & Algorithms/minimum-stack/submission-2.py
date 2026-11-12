class MinStack:
    def __init__(self):
        self.min_stack = []
        self.stack = []

    def push(self, val: int) -> None:
        if not self.min_stack or val <= self.min_stack[-1]:
            self.min_stack.append(val)
        self.stack.append(val)

    def pop(self) -> None:
        elem = self.stack.pop()
        if elem == self.min_stack[-1]:
            self.min_stack.pop()

    def top(self) -> int:
        return self.stack[-1]

    def getMin(self) -> int:
        if self.min_stack:
            return self.min_stack[-1]
        elif self.stack:
            return self.stack[-1]
        else:
            return None