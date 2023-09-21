from typing import List

class Solution:
    def threeSum(self, nums: List[int]) -> List[List[int]]:
        nums.sort()
        res = []
        for i in range(len(nums)):
            if nums[i] > 0: break
            if i == 0 or nums[i-1] != nums[i]:
                self.twoSum(nums, i, res)
        return res

    def twoSum(self, nums: List[int], i: int, res: List[List[int]]):
        lo, hi = i+1, len(nums)-1
        while lo < hi:
            sum = nums[i] + nums[lo] + nums[hi]
            if sum < 0 or (lo > i+1 and nums[lo] == nums[lo-1]):
                lo += 1
            elif sum > 0 or (hi < len(nums)-1 and nums[hi] == nums[hi+1]):
                hi -= 1
            else:
                res.append([nums[i], nums[lo], nums[hi]])
                lo += 1
                hi -= 1

if __name__ == "__main__":
    nums = [-1,0,1,2,-1,-4]
    sol = Solution()
    print(sol.threeSum(nums))
