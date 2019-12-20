# -*- coding: utf-8 -*-
# 5. 最长回文子串

def longestPalindrome(s):
    """
    :type s: str
    :rtype: str
    """
    ch_dic = {}
    for index, ch in enumerate(s):
        if ch_dic.get(ch) is None:
            ch_dic[ch] = [index]
        elif len(ch_dic[ch]) < 2:
            ch_dic[ch].append(index)
        else:
            ch_dic[ch][1] = index
    tmp = [0, 0]
    for key in ch_dic:
        if len(ch_dic[key]) < 2:
            continue
        if ch_dic[key][1] - ch_dic[key][0] > tmp[1] - tmp[0]:
            tmp[0] = ch_dic[key][0]
            tmp[1] = ch_dic[key][1]
    return s[tmp[0]:tmp[1]+1]



print longestPalindrome("babad")


