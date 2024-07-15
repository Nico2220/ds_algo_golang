package twopointers

import "fmt"


func SearchMatrxiTwoOptimalSolution(matrix [][]int, target int) bool{
	top := 0
    right := len(matrix[0]) - 1
    for top < len(matrix) && right >= 0 {
        v := matrix[top][right]
		fmt.Println(v, target)
        if v == target {
            return true
        }

        if v < target {
            top++
        }else {
            right--
        }
    }

    return false
}

func searchMatrix2(matrix [][]int, target int) bool {
    for i:=0; i < len(matrix); i++ {
        arr := matrix[i]

        if arr[0] <= target && target <= arr[len(arr) - 1] {
            if bs(arr, target) {
                return true
            }
        }
    }

    return false
}

func bs(arr []int, target int) bool {
    low, hight := 0, len(arr) - 1
    for low <= hight {
        mid := (low + hight) / 2
        if arr[mid] == target {
            return true
        } 

        if arr[mid] < target {
            low = mid + 1
        }else {
            hight = mid - 1
        }
    }

    return false
}