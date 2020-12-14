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
	data, err := ioutil.ReadFile("input")
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

	var divisors []int
	var remainders []int
	for i, s := range strings.Split(busesStr, ",") {
		num, err := strconv.Atoi(string(s))
		if err == nil {
			divisors = append(divisors, num)
			remainders = append(remainders, i)
		}
	}
	fmt.Println("Part2:", solveCRT(divisors, remainders))
	fmt.Println("Part2:", solveCRT2(divisors, remainders))
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

/* Solve a variation of the Chinese Remainder theorem

   a_i   divisors
x     mod  7 = 0
x + 1 mod 13 = 0
x + 4 mod 59 = 0
x + 6 mod 31 = 0
x + 7 mod 19 = 0

Iteratively solves a system of congruences, by finding searching x that
satisfies the first one, two, ... congruences. The step size is the product
of the divisors of the congruences that have been satisfied so far. Use the
fact that all divisors are prime.

Same as https://github.com/UnicycleBloke/aoc2020/blob/main/day13/day13.py#L89
*/
func solveCRT(divisors []int, as []int) int {
	x, step := 1, 1
	for i, d := range divisors {
		for (x+as[i])%d != 0 {
			x += step
		}
		step *= d
	}
	return x
}

/*
Solution based on Chinese Remainder theorem.
https://www.youtube.com/watch?v=ru7mWZJlRQg
https://github.com/UnicycleBloke/aoc2020/blob/main/day13/day13.py#L64
*/
func solveCRT2(divisors []int, as []int) int {
	M := 1 // product of all divisors
	for _, d := range divisors {
		M *= d
	}

	var sum int
	for i, d := range divisors {
		m := M / d
		prefactor := 0
		for (prefactor*m+as[i])%d != 0 {
			prefactor++
		}
		sum += prefactor * m
	}
	return sum % M
}
