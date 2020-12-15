package main

import (
	"fmt"
)

func main() {
	nums := []int{2, 20, 0, 4, 1, 17}
	for turn := len(nums) + 1; turn <= 2020; turn++ {
		lastNum := nums[len(nums)-1]
		prevTurn := findLast(nums[:len(nums)-1], lastNum) + 1
		if prevTurn == 0 {
			nums = append(nums, 0)
		} else {
			nums = append(nums, (turn-1)-(prevTurn))
		}
		//fmt.Println(nums)
	}
	fmt.Println("Part1:", nums[len(nums)-1])

	//################

	nums = []int{2, 20, 0, 4, 1, 17}
	lastTurn := map[int]int{}
	for i, elem := range nums {
		lastTurn[elem] = i + 1
	}
	var nextNum int
	lastNum := nums[len(nums)-1]
	for turn := len(nums) + 1; turn <= 30000000; turn++ {
		prevTurn, ok := lastTurn[lastNum]
		if !ok {
			nextNum = 0
		} else {
			nextNum = (turn - 1) - prevTurn
		}
		lastTurn[lastNum] = turn - 1
		lastNum = nextNum
	}
	fmt.Println("Part2:", nextNum)
}

// Find highest index of num in nums
func findLast(nums []int, num int) int {
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == num {
			return i
		}
	}
	return -1
}
