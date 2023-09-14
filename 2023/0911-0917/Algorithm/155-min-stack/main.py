class MinStack:

    def __init__(self):
        """
        initialize your data structure here.
        """
        self.stack = []
        self.min_val = float('inf')



    def push(self, val: int) -> None:
        self.stack.append(val)
        self.min_val = min(self.min_val, val)



    def pop(self) -> None:
        self.stack.pop()
        self.min_val = min(self.stack) if self.stack else float('inf')


    def top(self) -> int:
        return self.stack[-1]


    def getMin(self) -> int:
        return self.min_val



# Your MinStack object will be instantiated and called as such:
# obj = MinStack()
# obj.push(val)
# obj.pop()
# param_3 = obj.top()
# param_4 = obj.getMin()