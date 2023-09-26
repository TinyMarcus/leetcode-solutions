func lengthOfLongestSubstring(s string) int {
    var maxLength int
    var start int
    characters := map[rune]bool{}
    runeString := []rune(s)

    for end, character := range runeString {
        if !characters[character] || characters[character] == false {
            characters[character] = true
            maxLength = getMax(maxLength, end - start + 1)
        } else {
            for characters[character] == true {
                characters[runeString[start]] = false
                start += 1
            }

            characters[character] = true
        }
    }

    return maxLength
}

func getMax(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}