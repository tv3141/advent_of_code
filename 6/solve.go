package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var records []string
	var record []string
	recordSeparator := ""
	for scanner.Scan() {
		line := scanner.Text()
		if line == recordSeparator {
			records = append(records, strings.Join(record, ""))
			record = nil
			continue
		}
		record = append(record, line)
	}
	records = append(records, strings.Join(record, "")) // last record is not followed by empty line
	fmt.Println(len(records))

	var uniqueLetterCount int
	var uniqueLetterCounts []int
	for _, groupVoting := range records {
		uniqueLetterCount = setSize(groupVoting)
		uniqueLetterCounts = append(uniqueLetterCounts, uniqueLetterCount)
		fmt.Println(groupVoting, uniqueLetterCount)
	}
	sum := 0
	for _, count := range uniqueLetterCounts {
		sum += count
	}
	fmt.Println("Part1:", sum)

}

func setSize(items string) int {
	var uniqueCount int
	mymap := map[rune]int{}
	for _, item := range items {
		mymap[item] = 0
	}
	uniqueCount = 0
	for range mymap {
		uniqueCount++
	}
	return uniqueCount
}
