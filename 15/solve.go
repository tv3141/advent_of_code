package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 20, 0, 4, 1, 17}
	for turn := len(nums) + 1; turn <= 202000; turn++ {
		lastNum := nums[len(nums)-1]
		prevTurn := findLast(nums[:len(nums)-1], lastNum) + 1
		if prevTurn == 0 {
			nums = append(nums, 0)
		} else {
			nums = append(nums, (turn-1)-(prevTurn))
		}
	}
	fmt.Println("Part1:", nums[len(nums)-1])
}

// Find highest index of num in nums
func findLast(nums []int, num int) int {
	var indexes []int
	for i, elem := range nums {
		if elem == num {
			indexes = append(indexes, i)
		}
	}
	if len(indexes) == 0 {
		return -1
	}
	return indexes[len(indexes)-1]
}
