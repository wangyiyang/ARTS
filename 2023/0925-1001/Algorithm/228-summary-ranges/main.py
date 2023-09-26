from typing import List

class Solution:
    def summaryRanges(self, nums: List[int]) -> List[str]:
        if not nums:
            return []
        res = []
        start = nums[0]
        for i in range(1, len(nums)):
            if nums[i] - nums[i-1] != 1:
                if start != nums[i-1]:
                    res.append(str(start) + "->" + str(nums[i-1]))
                else:
                    res.append(str(start))
                start = nums[i]
        if start != nums[-1]:
            res.append(str(start) + "->" + str(nums[-1]))
        else:
            res.append(str(start))
        return res

if __name__ == "__main__":
    nums = [0,2,3,4,6,8,9]
    sol = Solution()
    print(sol.summaryRanges(nums))
