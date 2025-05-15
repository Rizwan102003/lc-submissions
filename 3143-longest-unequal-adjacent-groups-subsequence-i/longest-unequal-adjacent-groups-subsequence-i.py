class Solution:
    def getLongestSubsequence(self, words: List[str], groups: List[int]) -> List[str]:
        return [words[0]] + [w for i, w in enumerate(words[1:], start=1)
            if groups[i] != groups[i-1]] if words else []