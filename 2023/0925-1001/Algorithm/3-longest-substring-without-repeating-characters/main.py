class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        if not s:
            return 0
        max_len = 0
        for i in range(len(s)):
            for j in range(i, len(s)):
                if self.is_unique(s, i, j):
                    max_len = max(max_len, j - i + 1)
        return max_len

    def is_unique(self, s, start, end):
        chars = set()
        for i in range(start, end + 1):
            if s[i] in chars:
                return False
            chars.add(s[i])
        return True

# main
if __name__ == '__main__':
    s = Solution()
    print(s.lengthOfLongestSubstring("abcabcbb"))
