package main

import (
	"fmt"
	"io/ioutil"
	"math"
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
	ins := []string{}
	for _, line := range strings.Split(string(data), "\n") {
		if line != "" {
			ins = append(ins, line)
		}
	}

	x := 0
	y := 0
	curDir := "E"
	dir := "E"
	dirs := "NESW"
	for _, in := range ins {
		d, _ := strconv.Atoi(string(in[1:]))
		if strings.Contains("NESW", string(in[0])) {
			dir = string(in[0])
		} else {
			if string(in[0]) == "F" {
				// do nothing
			} else if string(in[0]) == "L" {
				dirChange := (d / 90) % 360
				curDir = string(dirs[(strings.Index(dirs, curDir)-dirChange+4)%4])
				continue
			} else if string(in[0]) == "R" {
				dirChange := (d / 90) % 360
				curDir = string(dirs[(strings.Index(dirs, curDir)+dirChange+4)%4])
				continue
			}
			dir = curDir
		}
		switch dir {
		case "N":
			y += d
		case "S":
			y -= d
		case "E":
			x += d
		case "W":
			x -= d
		}
	}
	fmt.Println("Part1:", math.Abs(float64(x))+math.Abs(float64(y)))

	// Part 2
	x = 0
	y = 0
	// waypoint
	wpX := 10
	wpY := 1
	for _, in := range ins {
		if strings.Contains("NESWLR", string(in[0])) { // move waypoint
			dir = string(in[0])
			d, _ := strconv.Atoi(string(in[1:]))
			switch dir {
			case "N":
				wpY += d
			case "S":
				wpY -= d
			case "E":
				wpX += d
			case "W":
				wpX -= d
			case "L":
				angle, _ := strconv.Atoi(string(in[1:]))
				for i := 0; i < (angle/90)%4; i++ {
					wpX, wpY = -wpY, wpX // 2d vector rotation
				}
			case "R":
				angle, _ := strconv.Atoi(string(in[1:]))
				for i := 0; i < (angle/90)%4; i++ {
					wpX, wpY = wpY, -wpX //2d vector rotation
				}
			}
		} else {
			n, _ := strconv.Atoi(string(in[1:]))
			x += n * wpX
			y += n * wpY
		}
	}
	fmt.Println("Part2:", math.Abs(float64(x))+math.Abs(float64(y)))
}
