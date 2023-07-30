class Solution:
    def maxDepth(self, s: str) -> int:
        ans, size = 0, 0
        for ch in s:
            if ch == '(':
                size += 1
                ans = max(ans, size)
            elif ch == ')':
                size -= 1
        return ans


if __name__=="__main__":
    a = Solution()
    print(a.maxDepth("(1+(2*3)+((8)/4))+1"))