class Solution:
    def canConstruct(self, ransomNote: str, magazine: str) -> bool:
        return all(ransomNote.count(i)<=magazine.count(i) for i in set(ransomNote))