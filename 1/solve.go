package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("inut")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var numbers []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		//fmt.Println(line)
		n, _ := strconv.Atoi(line)
		numbers = append(numbers, n)
	}
	//fmt.Println(numbers)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Part 1")
	for i1, num1 := range numbers {
		for _, num2 := range numbers[i1:] {
			if num1+num2 == 2020 {
				fmt.Println(num1, num2, num1*num2)
			}
		}
	}

	fmt.Println("Part 2")
	for i1, num1 := range numbers {
		for i2, num2 := range numbers[i1:] {
			for _, num3 := range numbers[i1+i2:] {
				if num1+num2+num3 == 2020 {
					fmt.Println(num1, num2, num3, num1*num2*num3)
				}
			}
		}
	}
}
