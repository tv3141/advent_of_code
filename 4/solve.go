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
			records = append(records, strings.Join(record, " "))
			record = nil
			continue
		}
		record = append(record, line)
	}
	records = append(records, strings.Join(record, " "))

	validCount := 0
	invalidCount := 0
	for _, record := range records {
		if strings.Contains(record, "byr:") &&
			strings.Contains(record, "iyr:") &&
			strings.Contains(record, "eyr:") &&
			strings.Contains(record, "hgt:") &&
			strings.Contains(record, "hcl:") &&
			strings.Contains(record, "ecl:") &&
			strings.Contains(record, "pid:") {
			validCount++
		} else {
			invalidCount++
			fmt.Println(record)
		}
	}
	fmt.Println("Part1:", validCount, invalidCount, len(records))
}
