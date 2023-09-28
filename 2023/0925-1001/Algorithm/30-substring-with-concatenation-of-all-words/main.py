from typing import List

class Solution:
    def findSubstring(self, s: str, words: List[str]) -> List[int]:
        if not s or not words: return []
        word_len = len(words[0])
        word_num = len(words)
        words_dict = {}
        for word in words:
            if word not in words_dict:
                words_dict[word] = 1
            else:
                words_dict[word] += 1
        res = []
        for i in range(len(s) - word_len * word_num + 1):
            seen = {}
            j = 0
            while j < word_num:
                word = s[i + j * word_len: i + (j + 1) * word_len]
                if word not in words_dict:
                    break
                if word not in seen:
                    seen[word] = 1
                else:
                    seen[word] += 1
                if seen[word] > words_dict[word]:
                    break
                j += 1
            if j == word_num:
                res.append(i)
        return res

if __name__ == '__main__':
    s = "barfoothefoobarman"
    words = ["foo","bar"]
    solution = Solution()
    print(solution.findSubstring(s, words))
