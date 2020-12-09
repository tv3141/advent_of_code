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
	data, err := ioutil.ReadFile("input")
	check(err)
	bags := make(map[string][]string)
	for _, line := range strings.Split(string(data), "\n") {
		if line != "" {
			bagColor, containedColors := parseRules(line)
			bags[bagColor] = containedColors
		}
	}
	reversedBags := reverseMap(bags)

	containsGold := make(map[string]bool)
	var queue []string
	queue = append(queue, reversedBags["shiny gold"]...)
	for len(queue) > 0 {
		nextColor := queue[0]
		containsGold[nextColor] = true
		queue = append(queue, reversedBags[nextColor]...)
		queue = queue[1:]
	}
	fmt.Println("Part 1:", len(containsGold))

	// ####################

	bags2 := make(map[string][]map[string]int)
	for _, line := range strings.Split(string(data), "\n") {
		if line != "" {
			bagColor, containedBags := parseRules2(line)
			bags2[bagColor] = containedBags
		}
	}

	var countBags func(color string) int
	countBags = func(color string) int {
		var count int
		if bags2[color] == nil {
			return 0
		}
		for _, bag := range bags2[color] {
			for color2, number := range bag {
				count += number * (countBags(color2) + 1)
			}
		}
		return count

	}

	fmt.Println("Part 2:", countBags("shiny gold"))

}

// parse rules ignoring number of bags
func parseRules(line string) (string, []string) {
	regexFullRule := regexp.MustCompile(`^(?P<bag>[a-z]+ [a-z]+) bags contain (?P<contained_bags>.+)\.$`)
	regexContainedColor := regexp.MustCompile(`\d+ (?P<contained_color>[a-z]+ [a-z]+) bag[s]?`)

	matches := regexFullRule.FindStringSubmatch(line)
	bagColor := matches[regexFullRule.SubexpIndex("bag")]

	containedBags := strings.Split(matches[regexFullRule.SubexpIndex("contained_bags")], ", ")
	var containedColors []string
	if containedBags[0] == "no other bags" {
		containedColors = nil
	} else {
		for _, cbag := range containedBags {
			containedColors = append(containedColors, regexContainedColor.FindStringSubmatch(cbag)[regexContainedColor.SubexpIndex("contained_color")])
		}
	}
	return bagColor, containedColors
}

// parse rules including number of bags
func parseRules2(line string) (string, []map[string]int) {
	regexFullRule := regexp.MustCompile(`^(?P<bag>[a-z]+ [a-z]+) bags contain (?P<contained_bags>.+)\.$`)
	regexContainedColor := regexp.MustCompile(`(?P<number>\d+) (?P<contained_color>[a-z]+ [a-z]+) bag[s]?`)

	matches := regexFullRule.FindStringSubmatch(line)
	bagColor := matches[regexFullRule.SubexpIndex("bag")]

	containedBagsStrings := strings.Split(matches[regexFullRule.SubexpIndex("contained_bags")], ", ")
	var containedBags []map[string]int
	if containedBagsStrings[0] == "no other bags" {
		containedBags = nil
	} else {
		var bag map[string]int
		for _, cbag := range containedBagsStrings {
			color := regexContainedColor.FindStringSubmatch(cbag)[regexContainedColor.SubexpIndex("contained_color")]
			number, _ := strconv.Atoi(regexContainedColor.FindStringSubmatch(cbag)[regexContainedColor.SubexpIndex("number")])
			bag = map[string]int{color: number}
			containedBags = append(containedBags, bag)
		}
	}
	return bagColor, containedBags
}

func reverseMap(m map[string][]string) map[string][]string {
	reversedMap := make(map[string][]string)
	for k, v := range m {
		for _, elem := range v {
			reversedMap[elem] = append(reversedMap[elem], string(k))
		}
	}
	return reversedMap
}
