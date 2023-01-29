class Solution:
    def isAnagram(s: str, t: str) -> bool:
        if len(s) != len(t):
            return False

        lettersS, lettersT = {}, {}

        for i in range(len(s)):
            lettersS[s[i]] = 1 + lettersS.get(s[i], 0)
            lettersT[t[i]] = 1 + lettersT.get(t[i], 0)

        for char in lettersS:
            if lettersS[char] != lettersT.get(char, 0):
                return False

        return True