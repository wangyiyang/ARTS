from typing import List


class Solution:
    def findMinArrowShots(self, points: List[List[int]]) -> int:
        points.sort(key=lambda x: x[1])

        arrows = 0
        prev_end = float('-inf')
        for start, end in points:
            if start > prev_end:
                arrows += 1
                prev_end = end
        return arrows

