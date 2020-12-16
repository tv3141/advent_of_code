package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Ticket struct {
	nums []int
}

type Rule struct {
	name   string
	range1 []int
	range2 []int
}

func main() {
	data, err := ioutil.ReadFile("input")
	check(err)

	rules := []Rule{}

	lines := strings.Split(string(data), "\n")
	for i := 0; i <= 19; i++ {
		re := regexp.MustCompile(`^(?P<name>[\w ]+): (?P<r1s>\d+)-(?P<r1e>\d+) or (?P<r2s>\d+)-(?P<r2e>\d+)$`)
		matches := re.FindStringSubmatch(lines[i])
		name := matches[re.SubexpIndex("name")]
		r1s, _ := strconv.Atoi(matches[re.SubexpIndex("r1s")])
		r1e, _ := strconv.Atoi(matches[re.SubexpIndex("r1e")])
		r2s, _ := strconv.Atoi(matches[re.SubexpIndex("r2s")])
		r2e, _ := strconv.Atoi(matches[re.SubexpIndex("r2e")])
		rule := Rule{name: name, range1: []int{r1s, r1e}, range2: []int{r2s, r2e}}
		rules = append(rules, rule)
	}

	myTicket := Ticket{}
	for _, i := range strings.Split(strings.Split(string(data), "\n")[22], ",") {
		num, _ := strconv.Atoi(string(i))
		myTicket.nums = append(myTicket.nums, num)
	}

	var otherTickets []Ticket
	for i := 25; i <= 260; i++ {
		ticket := Ticket{}
		for _, i := range strings.Split(strings.Split(string(data), "\n")[i], ",") {
			num, _ := strconv.Atoi(string(i))
			ticket.nums = append(ticket.nums, num)
		}
		otherTickets = append(otherTickets, ticket)
	}

	//fmt.Println(rules, myTicket, otherTickets)

	invalidSum := 0
	for _, t := range otherTickets {
		for _, n := range t.nums {
			valid := false
			for _, r := range rules {
				if (r.range1[0] <= n && n <= r.range1[1]) || (r.range2[0] <= n && n <= r.range2[1]) {
					valid = true
					break
				}
			}
			if !valid {
				invalidSum += n
			}
		}

	}
	fmt.Println("Part1:", invalidSum)

	// Part 2
	var validOtherTickets []Ticket
	for _, t := range otherTickets {
		validTicket := true
		for _, n := range t.nums {
			valid := false
			for _, r := range rules {
				if (r.range1[0] <= n && n <= r.range1[1]) || (r.range2[0] <= n && n <= r.range2[1]) {
					valid = true
					break
				} else {
					//rule invalid
				}
			}
			if !valid {
				validTicket = false
			}
		}
		if validTicket {
			validOtherTickets = append(validOtherTickets, t)
		}

	}
	//fmt.Println(len(validOtherTickets))

	// store number as columns
	cols := [][]int{}
	for i := 0; i < len(validOtherTickets[0].nums); i++ {
		col := []int{}
		for _, t := range validOtherTickets {
			//fmt.Println(t.nums[0])
			col = append(col, t.nums[i])
		}
		cols = append(cols, col)
	}

	// rules columns
	possibleFields := make(map[int][]string)
	for colNum, col := range cols {
		for _, r := range rules {
			valid := true
			for _, n := range col {
				if (r.range1[0] <= n && n <= r.range1[1]) || (r.range2[0] <= n && n <= r.range2[1]) {
					// valid
				} else {
					valid = false
					break
				}
			}
			if valid {
				possibleFields[colNum] = append(possibleFields[colNum], r.name)
			}
		}
	}

	// constraint propagation
	fields := make(map[string]int)
	for len(possibleFields) > 0 {
		for k, v := range possibleFields {
			if len(v) == 1 {
				fields[v[0]] = k
				delete(possibleFields, k)
				for k := range possibleFields {
					possibleFields[k] = remove(possibleFields[k], v[0])
				}
			}
		}
	}
	//fmt.Println(fields)

	product := 1
	for k, v := range fields {
		if strings.Contains(k, "departure") {
			product *= myTicket.nums[v]
		}
	}

	fmt.Println("Part2:", product)
}

func remove(slice []string, s string) []string {
	result := []string{}
	for _, elem := range slice {
		if elem != s {
			result = append(result, elem)
		}
	}
	return result
}
