func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    letters := [26]int{}

    for _, val := range s {
        letters[val - 'a'] += 1
    }

    for _, val := range t {
        letters[val - 'a'] -= 1
    }

    for _, val := range letters {
        if val != 0 {
            return false
        }
    }

    return true
}