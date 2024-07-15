package array

import (
	"math"
)

func Insert(intervals [][]int, newInterval []int) [][]int {
    result := [][]int{}

    i := 0
    n := len(intervals)
    for i < n && intervals[i][1] < newInterval[0] { 
        result = append(result, intervals[i])
        i++
    }

    for i < n && intervals[i][0] <= newInterval[1] {
        newInterval[0] = int(math.Min(float64(intervals[i][0]), float64(newInterval[0])))
        newInterval[1] = int(math.Max(float64(intervals[i][1]), float64(newInterval[1])))
        i++
    }

    result = append(result, newInterval)

    for i < n {
        result = append(result, intervals[i])
        i++
    }

    return result
}