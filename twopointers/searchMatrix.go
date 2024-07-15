package twopointers

func searchMatrix(matrix [][]int, target int) bool {
    m, n := len(matrix), len(matrix[0])
    low, hight := 0, (m * n) - 1

    for low <= hight {
        mid := low + (hight - low) / 2
        row := mid / m 
		col := mid % m

        if matrix[row][col] == target {
            return true
        }

        if matrix[row][col] < target {
            low = mid + 1
        }else {
            hight = mid - 1
        }
    }

    return false
}


func SearchMatrix(matrix [][]int, target int) bool {
    m := len(matrix)
    n := len(matrix[0])
    left, right := 0, m*n-1

    for left <= right {
        mid := left + (right-left)/2
        mid_val := matrix[mid/n][mid%n]

        if mid_val == target {
            return true
        } else if mid_val < target {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return false
}