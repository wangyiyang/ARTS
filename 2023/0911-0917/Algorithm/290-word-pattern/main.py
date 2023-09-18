
class Solution:

    def wordPattern(self, pattern: str, s: str) -> bool:
        words = s.split()
        if len(words) != len(pattern):
            return False

        pattern_dict = {}
        word_dict = {}
        for i in range(len(pattern)):
            if pattern[i] not in pattern_dict:
                pattern_dict[pattern[i]] = words[i]
            elif pattern_dict[pattern[i]] != words[i]:
                return False

            if words[i] not in word_dict:
                word_dict[words[i]] = pattern[i]
            elif word_dict[words[i]] != pattern[i]:
                return False

        return True


if __name__ == "__main__":
    sol = Solution()
    pattern = "abba"
    s = "dog cat cat dog"
    print(sol.wordPattern(pattern, s))
    pattern = "abba"
    s = "dog cat cat fish"
    print(sol.wordPattern(pattern, s))
    pattern = "aaaa"
    s = "dog cat cat dog"
    print(sol.wordPattern(pattern, s))
    pattern = "abba"
    s = "dog dog dog dog"
    print(sol.wordPattern(pattern, s))
