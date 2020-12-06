package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Part 1
	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		pw, policy := parse(line)
		if checkPolicy(pw, policy) {
			count++
		}
	}
	fmt.Println("Part 1:", count)

	// Part 2
	file, err = os.Open("input")
	scanner = bufio.NewScanner(file)
	count = 0
	for scanner.Scan() {
		line := scanner.Text()
		pw, policy := parse2(line)
		if checkPolicy2(pw, policy) {
			count++
		}
	}
	fmt.Println("Part 2:", count)
}

type policy struct {
	min    int
	max    int
	letter string
}

func parse(s string) (string, policy) {
	fields := strings.Split(s, " ")
	pw := fields[2]
	letter := string(fields[1][0])
	min, _ := strconv.Atoi(strings.Split(fields[0], "-")[0])
	max, _ := strconv.Atoi(strings.Split(fields[0], "-")[1])
	policy := policy{min, max, letter}
	return pw, policy
}

func checkPolicy(pw string, policy policy) bool {
	letterCount := 0
	for _, char := range pw {
		if string(char) == policy.letter {
			letterCount++
		}
	}
	if policy.min <= letterCount && letterCount <= policy.max {
		return true
	}
	return false
}

type policy2 struct {
	pos1   int
	pos2   int
	letter string
}

func parse2(s string) (string, policy2) {
	re := regexp.MustCompile(`^(?P<pos1>\d+)-(?P<pos2>\d+)\s(?P<letter>\S):\s(?P<pw>\S+)$`)
	matches := re.FindStringSubmatch(s)

	pw := matches[re.SubexpIndex("pw")]
	pos1, _ := strconv.Atoi(matches[re.SubexpIndex("pos1")])
	pos2, _ := strconv.Atoi(matches[re.SubexpIndex("pos2")])
	letter := matches[re.SubexpIndex("letter")]

	policy := policy2{pos1, pos2, letter}
	return pw, policy
}

func checkPolicy2(pw string, policy policy2) bool {
	return (string(pw[policy.pos1-1]) == policy.letter) != (string(pw[policy.pos2-1]) == policy.letter)
}
