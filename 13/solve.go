package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	data, err := ioutil.ReadFile("test_input")
	check(err)
	t, _ := strconv.Atoi(strings.Split(string(data), "\n")[0])
	busesStr := strings.Split(string(data), "\n")[1]
	var buses []int
	for _, s := range strings.Split(busesStr, ",") {
		num, err := strconv.Atoi(string(s))
		if err == nil {
			buses = append(buses, num)
		}
	}
	nextBus, earliestTime := nextBus(buses, t)

	fmt.Println("Part1:", nextBus*(earliestTime-t))

	fmt.Println("Part2:")
	// Chinese remainder theorem
}

func nextBus(buses []int, t int) (int, int) {
	var earliestTime int
	var nextBus int
	for i, b := range buses {
		if i == 0 {
			earliestTime = int(math.Ceil(float64(t)/float64(b))) * b
			nextBus = b
		}
		nextTime := int(math.Ceil(float64(t)/float64(b))) * b
		if nextTime < earliestTime {
			earliestTime = nextTime
			nextBus = b
		}

	}
	return nextBus, earliestTime
}
