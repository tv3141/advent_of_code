package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type instruction struct {
	op  string
	arg int
}

func main() {
	data, err := ioutil.ReadFile("input")
	check(err)
	var numbers []int
	for _, line := range strings.Split(string(data), "\n") {
		if line != "" {
			num, _ := strconv.Atoi(line)
			numbers = append(numbers, num)
		}
	}
	windowSize := 25 // 5 for test_input
	firstInvalid := checkNums(numbers, windowSize)
	fmt.Println("Part1:", firstInvalid)

	// Part 2 ###########
	for i1 := range numbers {
		for i2 := i1 + 1; i2 < len(numbers); i2++ {
			if sum(numbers[i1:i2]) == firstInvalid {
				fmt.Println("Part2:", max(numbers[i1:i2])+min(numbers[i1:i2]))
				return
			}
		}
	}
}

func checkNums(nums []int, windowSize int) int {
	var invalidNum int
	for i := windowSize; i < len(nums); i++ {
		valid := verify(nums[i-windowSize:i], nums[i])
		if !valid {
			invalidNum = nums[i]
		}
	}
	return invalidNum
}

func verify(nums []int, toVerify int) bool {
	isSum := false
	for i1, num1 := range nums {
		for _, num2 := range nums[i1:] {
			if num1+num2 == toVerify {
				isSum = true
			}
		}
	}
	return isSum
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func max(array []int) int {
	result := array[0]
	for _, v := range array {
		if v > result {
			result = v
		}
	}
	return result
}

func min(array []int) int {
	result := array[0]
	for _, v := range array {
		if v < result {
			result = v
		}
	}
	return result
}
