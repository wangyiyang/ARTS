from typing import List


class Solution:
    def lengthOfLIS(self, nums: List[int]) -> int:
        # 动态规划
        # 时间复杂度：O(n^2)
        # 空间复杂度：O(n)
        # dp[i] 表示以 nums[i] 结尾的最长上升子序列的长度
        # dp[i] = max(dp[j]) + 1, 0 <= j < i and nums[j] < nums[i]
        # dp[0] = 1
        # return max(dp)
        if not nums:
            return 0

        dp = [1] * len(nums)
        for i in range(len(nums)):
            for j in range(i):
                if nums[j] < nums[i]:
                    # dp[i] = max(dp[j]) + 1, 0 <= j < i and nums[j] < nums[i]
                    dp[i] = max(dp[i], dp[j] + 1)

        return max(dp)


if __name__ == '__main__':
    nums = [10, 9, 2, 5, 3, 7, 101, 18]  # 4
    # nums = [0, 1, 0, 3, 2, 3]  # 4
    # nums = [7, 7, 7, 7, 7, 7, 7]  # 1
    # nums = [4, 10, 4, 3, 8, 9]  # 3
    # nums = [1, 3, 6, 7, 9, 4, 10, 5, 6]  # 6
    solution = Solution()
    print(solution.lengthOfLIS(nums))
