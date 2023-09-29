class Solution:
    def minWindow(self, s: str, t: str) -> str:
        # 1. init
        left = 0
        right = 0
        min_len = float('inf')
        min_left = 0
        min_right = 0
        window = {}
        needs = {}
        match = 0

        for c in t:
            needs[c] = needs.get(c, 0) + 1

        # 2. sliding window
        while right < len(s):
            c1 = s[right]
            if c1 in needs:
                window[c1] = window.get(c1, 0) + 1

                if window[c1] == needs[c1]:
                    match += 1

            right += 1

            while match == len(needs):
                if right - left < min_len:
                    min_len = right - left
                    min_left = left
                    min_right = right

                c2 = s[left]
                if c2 in needs:
                    window[c2] -= 1

                    if window[c2] < needs[c2]:
                        match -= 1

                left += 1

        return s[min_left:min_right] if min_len != float('inf') else ''