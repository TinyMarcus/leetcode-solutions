# Longest substring without repeating characters (3): https://leetcode.com/problems/longest-substring-without-repeating-characters/

class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        maximum = 0
        start = 0
        characters = dict()

        for end, character in enumerate(s):
            if character not in characters.keys() or characters[character] == False:
                characters[character] = True
                maximum = max(maximum, end - start + 1)
            else:
                while characters[character]:
                    characters[s[start]] = False
                    start += 1
            
            characters[character] = True
        
        return maximum