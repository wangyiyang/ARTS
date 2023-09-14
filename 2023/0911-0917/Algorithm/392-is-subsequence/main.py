class Solution(object):
    def isSubsequence(self, s, t):
        """
        :type s: str
        :type t: str
        :rtype: bool
        """
        s_index = 0
        for i in range(len(t)):
            if s_index == len(s):
                return True
            if s[s_index] == t[i]:
                s_index += 1
        return s_index == len(s)

if __name__ == '__main__':
    s = Solution()
    print(s.isSubsequence('abc', 'ahbgdc'))
    print(s.isSubsequence('axc', 'ahbgdc'))
    print(s.isSubsequence('acb', 'ahbgdc'))
    print(s.isSubsequence('abc', 'ahbgdca'))
    print(s.isSubsequence('abc', 'ahbgd'))
