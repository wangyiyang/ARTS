from typing import List

class Solution:
    def romanToInt(self, s: str) -> int:
        res = 0
        n = len(s)
        i = 0
        while i < n:
            if s[i] == 'I':
                if i + 1 < n and s[i + 1] == 'V':
                    res += 4
                    i += 2
                elif i + 1 < n and s[i + 1] == 'X':
                    res += 9
                    i += 2
                else:
                    res += 1
                    i += 1
            elif s[i] == 'V':
                res += 5
                i += 1
            elif s[i] == 'X':
                if i + 1 < n and s[i + 1] == 'L':
                    res += 40
                    i += 2
                elif i + 1 < n and s[i + 1] == 'C':
                    res += 90
                    i += 2
                else:
                    res += 10
                    i += 1
            elif s[i] == 'L':
                res += 50
                i += 1
            elif s[i] == 'C':
                if i + 1 < n and s[i + 1] == 'D':
                    res += 400
                    i += 2
                elif i + 1 < n and s[i + 1] == 'M':
                    res += 900
                    i += 2
                else:
                    res += 100
                    i += 1
            elif s[i] == 'D':
                res += 500
                i += 1
            elif s[i] == 'M':
                res += 1000
                i += 1
        return res
    
if __name__ == '__main__':
    s = "MCMXCIV"
    solution = Solution()
    print(solution.romanToInt(s))
    