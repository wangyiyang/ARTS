# -*- coding: utf-8 -*-
# 5. 最长回文子串

# def longestPalindrome(s):
#     """
#     :type s: str
#     :rtype: str
#     """
#     ch_dic = {}
#     for index, ch in enumerate(s):
#         if ch_dic.get(ch) is None:
#             ch_dic[ch] = [index]
#         elif len(ch_dic[ch]) < 2:
#             ch_dic[ch].append(index)
#         else:
#             ch_dic[ch][1] = index
#     tmp = [0, 0]
#     for key in ch_dic:
#         if len(ch_dic[key]) < 2:
#             continue
#         if ch_dic[key][1] - ch_dic[key][0] > tmp[1] - tmp[0]:
#             tmp[0] = ch_dic[key][0]
#             tmp[1] = ch_dic[key][1]
#     return s[tmp[0]:tmp[1]+1]

def longestPalindrome(s):
    # 特判
    size = len(s)
    if size < 2:
        return s

    max_len = 1
    res = s[0]

    # 枚举所有长度大于等于 2 的子串
    for i in range(size - 1):
        for j in range(i + 1, size):
            if j - i + 1 > max_len and __valid(s, i, j):
                max_len = j - i + 1
                res = s[i:j + 1]
    return res

def __valid(s, left, right):
    # 验证子串 s[left, right] 是否为回文串
    while left < right:
        if s[left] != s[right]:
            return False
        left += 1
        right -= 1
    return True


print longestPalindrome("babad")
# print longestPalindrome("abcda")


