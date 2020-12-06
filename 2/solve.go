package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		pw, policy := parse(line)
		if checkPolicy(pw, policy) {
			count++
		}
	}
	fmt.Println("Part 1:", count)

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
