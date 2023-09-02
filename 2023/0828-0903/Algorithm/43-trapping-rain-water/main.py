from types import List

class Solution:
    def trap(self, height: List[int]) -> int:
        n = len(height)
        left = [0] * n
        right = [0] * n
        for i in range(1, n):
            left[i] = max(left[i - 1], height[i - 1])
        for i in range(n - 2, -1, -1):
            right[i] = max(right[i + 1], height[i + 1])
        res = 0
        for i in range(n):
            res += max(0, min(left[i], right[i]) - height[i])
        return res

if __name__ == '__main__':
    height = [0,1,0,2,1,0,1,3,2,1,2,1]
    solution = Solution()
    print(solution.trap(height))