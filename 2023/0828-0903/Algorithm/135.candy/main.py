from typing import List

class Solution:
    def candy(self, ratings: List[int]) -> int:
        n = len(ratings)
        left = [1] * n
        right = [1] * n
        for i in range(1, n):
            if ratings[i] > ratings[i - 1]:
                left[i] = left[i - 1] + 1
        for i in range(n - 2, -1, -1):
            if ratings[i] > ratings[i + 1]:
                right[i] = right[i + 1] + 1
        res = 0
        for i in range(n):
            res += max(left[i], right[i])
        return res

if __name__ == '__main__':
    ratings = [1,0,2]
    solution = Solution()
    print(solution.candy(ratings))