package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	data, err := ioutil.ReadFile("input")
	check(err)
	numbers := []int{0}
	for _, line := range strings.Split(string(data), "\n") {
		if line != "" {
			n, _ := strconv.Atoi(line)
			numbers = append(numbers, n)
		}
	}
	sort.Ints(numbers)
	ones := 0
	threes := 1 // last adapter is always +3
	for i := 0; i < (len(numbers) - 1); i++ {
		diff := numbers[i+1] - numbers[i]
		if diff == 3 {
			threes++
		} else if diff == 1 {
			ones++
		} else {
			fmt.Println("Unexpected diff:", diff)
		}
	}
	fmt.Println("Part1:", ones*threes)
}
