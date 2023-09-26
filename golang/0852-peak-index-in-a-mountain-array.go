func peakIndexInMountainArray(arr []int) int {
    start := 0
    end := len(arr) - 1

    for start <= end {
        middle := (start + end) / 2

        if arr[middle] < arr[middle + 1] {
            start = middle + 1
        } else {
            end = middle - 1
        }
    }
    
    return start
}