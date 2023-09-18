from typing import List


class Solution:
    def findMinArrowShots(self, points: List[List[int]]) -> int:
            """
            Given a list of intervals, returns the minimum number of arrows that must be shot to burst all balloons.

            Args:
                points (List[List[int]]): A list of intervals, where each interval is represented as a list of two integers [start, end].

            Returns:
                int: The minimum number of arrows that must be shot to burst all balloons.
            """
            points.sort(key=lambda x: x[1])

            arrows = 0
            prev_end = float('-inf')
            for start, end in points:
                if start > prev_end:
                    arrows += 1
                    prev_end = end
            return arrows

