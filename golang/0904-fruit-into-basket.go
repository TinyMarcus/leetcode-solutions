func totalFruit(fruits []int) int {
    var maxCount int 
    var start int
    baskets := map[int]bool{}

    for end, treeType := range fruits {
        if len(baskets) < 2 && !baskets[treeType] {
            baskets[treeType] = true
            maxCount = getMax(maxCount, end - start + 1)
        } else if baskets[treeType] {
            maxCount = getMax(maxCount, end - start + 1)
        } else {
            baskets = map[int]bool{}
            baskets[fruits[end - 1]] = true
            baskets[treeType] = true
            start = end - 1

            for fruits[start] == fruits[start - 1] {
                start--
            }

            maxCount = getMax(maxCount, end - start + 1)
        }
    }

    return maxCount
}

func getMax(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}