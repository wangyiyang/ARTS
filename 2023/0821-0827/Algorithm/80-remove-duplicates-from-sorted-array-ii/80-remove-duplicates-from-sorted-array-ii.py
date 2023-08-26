from typing import List

class Solution:
    def removeDuplicates(self, nums: List[int]) -> int:
        """
        Given a sorted array nums, remove the duplicates in-place such that duplicates appeared at most twice and return the new length.

        Args:
        - nums: a list of integers representing the sorted array

        Returns:
        - an integer representing the new length of the array after removing duplicates
        """
        i = 0
        for n in nums:
            if i < 2 or n > nums[i-2]:
                nums[i] = n
                i += 1
        return i
