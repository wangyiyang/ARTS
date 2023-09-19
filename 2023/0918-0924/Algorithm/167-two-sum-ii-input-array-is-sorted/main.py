from typing import List

# 给你一个下标从 1 开始的整数数组 numbers ，该数组已按 非递减顺序排列  ，请你从数组中找出满足相加之和等于目标数 target 的两个数。如果设这两个数分别是 numbers[index1] 和 numbers[index2] ，则 1 <= index1 < index2 <= numbers.length 。


class Solution:
    def twoSum(self, numbers: List[int], target: int) -> List[int]:
        # 二分查找
        # 时间复杂度：O(nlogn)
        # 空间复杂度：O(1)
        for i in range(len(numbers)):
            low, high = i + 1, len(numbers) - 1
            while low <= high:
                mid = (low + high) // 2
                if numbers[mid] == target - numbers[i]:
                    return [i + 1, mid + 1]
                elif numbers[mid] > target - numbers[i]:
                    high = mid - 1
                else:
                    low = mid + 1

        return [-1, -1]