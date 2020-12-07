package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

	var seatIDs []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		binarise := func(r rune) rune {
			var res rune
			switch r {
			case 'F':
				res = '0'
			case 'B':
				res = '1'
			case 'L':
				res = '0'
			case 'R':
				res = '1'
			default:
				log.Fatal("Unknown seat encoding")
			}
			return res
		}

		row, _ := strconv.ParseInt(strings.Map(binarise, line[0:7]), 2, 32)
		col, _ := strconv.ParseInt(strings.Map(binarise, line[7:10]), 2, 32)

		fmt.Println(line, "Row:", row, "Col:", col)
		seatIDs = append(seatIDs, seatID(int(row), int(col)))
	}
	maxSeatID := 0
	for _, sID := range seatIDs {
		if sID > maxSeatID {
			maxSeatID = sID
		}
	}
	fmt.Println("Part1:", maxSeatID)

	fmt.Println("Part2:", mySeatID(seatIDs))
}

func seatID(row, col int) int {
	return row*8 + col
}

func mySeatID(seatIDs []int) int {
	// Find the gap in seat IDs
	sort.Ints(seatIDs)
	var mySeatID int
	for i, id := range seatIDs {
		if seatIDs[i+1]-id == 2 {
			mySeatID = seatIDs[i] + 1
			break
		}
	}
	return mySeatID
}
