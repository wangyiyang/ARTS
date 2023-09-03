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

# 这段代码定义了一个名为Solution的类，其中包含一个名为trap的方法。该方法接受一个整数列表height作为参数，并返回一个整数值。该方法的目的是计算给定高度列表中的雨水总量。

# 该方法使用三个列表：left，right和res。left和right列表用于存储每个位置左侧和右侧的最大高度。res列表用于存储最终结果，即雨水总量。

# 该方法首先计算列表height的长度，并创建两个长度为n的空列表left和right。然后，使用两个循环分别计算每个位置左侧和右侧的最大高度。最后，使用第三个循环计算每个位置上的雨水量，并将其添加到res中。最后，返回res作为结果。

# 该算法的时间复杂度为$O(n)$，其中$n$是列表height的长度。 ▌

if __name__ == '__main__':
    height = [0,1,0,2,1,0,1,3,2,1,2,1]
    solution = Solution()
    print(solution.trap(height))