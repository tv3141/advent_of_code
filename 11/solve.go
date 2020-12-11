package main

import (
	"fmt"
	"io/ioutil"
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
	seatGrid := [][]string{}
	newGrid := [][]string{}
	for _, line := range strings.Split(string(data), "\n") {
		if line != "" {
			seatGrid = append(seatGrid, strings.Split(line, ""))
			newGrid = append(newGrid, strings.Split(line, ""))
		}
	}
	//printGrid(seatGrid)

	count := 0
	diff := 1
	for diff != 0 {
		for x := 0; x < len(seatGrid[0]); x++ {
			for y := 0; y < len(seatGrid); y++ {
				if seatGrid[y][x] != "." {
					neighbors := []string{}
					if x < len(seatGrid[0])-1 {
						neighbors = append(neighbors, string(seatGrid[y][x+1]))
					}
					if x > 0 {
						neighbors = append(neighbors, string(seatGrid[y][x-1]))
					}
					if x > 0 && y < len(seatGrid)-1 {
						neighbors = append(neighbors, string(seatGrid[y+1][x-1]))
					}
					if y < len(seatGrid)-1 {
						neighbors = append(neighbors, string(seatGrid[y+1][x]))
					}
					if x < len(seatGrid[0])-1 && y < len(seatGrid)-1 {
						neighbors = append(neighbors, string(seatGrid[y+1][x+1]))
					}
					if x > 0 && y > 0 {
						neighbors = append(neighbors, string(seatGrid[y-1][x-1]))
					}
					if y > 0 {
						neighbors = append(neighbors, string(seatGrid[y-1][x]))
					}
					if x < len(seatGrid[0])-1 && y > 0 {
						neighbors = append(neighbors, string(seatGrid[y-1][x+1]))
					}
					neighborsStr := strings.Join(neighbors, "")
					if seatGrid[y][x] == "L" && countChar(neighborsStr, "#") == 0 {
						newGrid[y][x] = "#"
					}
					if seatGrid[y][x] == "#" && countChar(neighborsStr, "#") >= 4 {
						newGrid[y][x] = "L"
					}
				}
			}
		}
		count = occupiedSeats(seatGrid)
		newCount := occupiedSeats(newGrid)
		diff = count - newCount
		for x := 0; x < len(seatGrid[0]); x++ {
			for y := 0; y < len(seatGrid); y++ {
				seatGrid[y][x] = newGrid[y][x]
			}
		}
	}
	//printGrid(seatGrid)
	fmt.Println("Part1:", count)
}

func countChar(s string, c string) int {
	return len(s) - len(strings.ReplaceAll(s, c, ""))
}

func occupiedSeats(grid [][]string) int {
	count := 0
	for x := 0; x < len(grid[0]); x++ {
		for y := 0; y < len(grid); y++ {
			if grid[y][x] == "#" {
				count++
			}
		}
	}
	return count
}

func printGrid(g [][]string) {
	for _, row := range g {
		fmt.Println(strings.Join(row, ""))
	}
	fmt.Println()
}
