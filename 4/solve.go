package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
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

	var sortedRecords []string
	for _, record := range records {
		splitRecord := strings.Split(record, " ")
		sort.Slice(splitRecord, func(i int, j int) bool {
			return strings.Split(string(splitRecord[i]), ":")[0] < strings.Split(string(splitRecord[j]), ":")[0]
		})
		sortedRecords = append(sortedRecords, strings.Join(splitRecord, " "))
	}

	var validPassports []string
	for _, record := range sortedRecords {
		if strings.Contains(record, "byr:") &&
			strings.Contains(record, "iyr:") &&
			strings.Contains(record, "eyr:") &&
			strings.Contains(record, "hgt:") &&
			strings.Contains(record, "hcl:") &&
			strings.Contains(record, "ecl:") &&
			strings.Contains(record, "pid:") {
			validPassports = append(validPassports, record)
		}
	}
	fmt.Println("Part1:", len(validPassports))

	// Part 2
	var validPassports2 []string
	for _, passport := range validPassports {
		invalid := false
		for _, field := range strings.Split(passport, " ") {
			fieldName := strings.Split(field, ":")[0]
			fieldValue := strings.Split(field, ":")[1]
			switch fieldName {
			case "byr":
				year, err := strconv.Atoi(fieldValue)
				if err != nil || year < 1920 || year > 2002 {
					invalid = true
				}
			case "iyr":
				year, err := strconv.Atoi(fieldValue)
				if err != nil || year < 2010 || year > 2020 {
					invalid = true
				}
			case "eyr":
				year, err := strconv.Atoi(fieldValue)
				if err != nil || year < 2020 || year > 2030 {
					invalid = true
				}
			case "hgt":
				re := regexp.MustCompile(`^\d+(in|cm)$`)
				if !re.MatchString(fieldValue) {
					invalid = true
				}

				re = regexp.MustCompile(`^(?P<height>\d+)in$`)
				matches := re.FindStringSubmatch(fieldValue)
				if len(matches) > 0 {
					heightInch, err := strconv.Atoi(matches[re.SubexpIndex("height")])
					if err != nil || heightInch < 59 || heightInch > 76 {
						invalid = true
					}
				}

				re = regexp.MustCompile(`^(?P<height>\d+)cm$`)
				matches = re.FindStringSubmatch(fieldValue)
				if len(matches) > 0 {
					heightCM, err := strconv.Atoi(matches[re.SubexpIndex("height")])
					if err != nil || heightCM < 150 || heightCM > 193 {
						invalid = true
					}
				}
			case "hcl":
				re := regexp.MustCompile(`^#[a-f0-9]{6}$`)
				if !re.MatchString(fieldValue) {
					invalid = true
				}
			case "ecl":
				re := regexp.MustCompile(`^(amb|blu|brn|gry|grn|hzl|oth)$`)
				if !re.MatchString(fieldValue) {
					invalid = true
				}
			case "pid":
				re := regexp.MustCompile(`^[0-9]{9}$`)
				if !re.MatchString(fieldValue) {
					invalid = true
				}
			}
		}
		if !invalid {
			validPassports2 = append(validPassports2, passport)
		}
	}
	fmt.Println("Part2:", len(validPassports2))
	/* 	for _, p := range validPassports2 {
		fmt.Println(p)
	} */
}
