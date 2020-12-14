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

func main() {
	data, err := ioutil.ReadFile("test_input")
	check(err)
	var mask string
	var mem = make(map[int64]int64)
	for _, line := range strings.Split(string(data), "\n") {
		if line != "" {
			if line[0:3] == "mas" {
				mask = strings.Split(line, " = ")[1]
			} else {
				re := regexp.MustCompile(`^mem\[(?P<address>\d+)\] = (?P<value>\d+)$`)
				matches := re.FindStringSubmatch(line)
				address, _ := strconv.Atoi(matches[re.SubexpIndex("address")])
				n, _ := strconv.Atoi(matches[re.SubexpIndex("value")])
				mem[int64(address)] = maskInt(int64(n), mask)
			}
		}
	}
	sum := int64(0)
	for _, v := range mem {
		sum += v
	}
	fmt.Println("Part1:", sum)

	fmt.Println("Part2:")
}

func maskInt(i int64, mask string) int64 {
	maskAND, _ := strconv.ParseInt(strings.Replace(mask, "X", "1", -1), 2, 64)
	maskOR, _ := strconv.ParseInt(strings.Replace(mask, "X", "0", -1), 2, 64)
	i &= maskAND
	i |= maskOR
	return i
}
