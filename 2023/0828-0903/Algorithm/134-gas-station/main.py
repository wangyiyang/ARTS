from typing import List


class Solution:
    def canCompleteCircuit(self, gas: List[int], cost: List[int]) -> int:
        # 1. 暴力法
        # for i in range(len(gas)):
        #     if gas[i] < cost[i]:
        #         continue
        #     else:
        #         if self.canComplete(gas, cost, i):
        #             return i
        # return -1

        # 2. 贪心算法
        # 1. 如果总油量小于总消耗，那么一定无法完成
        # 2. 如果某个站点的油量小于消耗，那么这个站点一定不是起点
        # 3. 如果从站点i出发，到达站点j时油量小于0，那么从i到j之间的任何一个站点都不可能是起点
        # 4. 如果从站点i出发，到达站点j时油量大于0，那么从i到j之间的任何一个站点都可能是起点
        # 5. 如果从站点i出发，到达终点时油量大于0，那么i就是起点
        if sum(gas) < sum(cost):
            return -1
        start = 0
        cur = 0
        for i in range(len(gas)):
            cur += gas[i] - cost[i]
            if cur < 0:
                start = i + 1
                cur = 0
        return start

        
    
if __name__ == '__main__':
    gas = [1,2,3,4,5]
    cost = [3,4,5,1,2]
    solution = Solution()
    print(solution.canCompleteCircuit(gas, cost))