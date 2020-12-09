package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type instruction struct {
	op  string
	arg int
}

func main() {
	data, err := ioutil.ReadFile("input")
	check(err)
	var instructions []instruction
	for _, line := range strings.Split(string(data), "\n") {
		if line != "" {
			inst := parseLine(line)
			//fmt.Println(line, inst)
			instructions = append(instructions, inst)
		}
	}
	acc, _ := execute(instructions)
	fmt.Println("Part1:", acc)

	// Part 2
	insCopy := make([]instruction, len(instructions))
	for i, ins := range instructions {
		copy(insCopy, instructions)
		if ins.op == "jmp" {
			insCopy[i] = instruction{"nop", 0}
		}
		acc, loopDetected := execute(insCopy)
		if !loopDetected {
			fmt.Println("Part2:", acc)
		}

	}
}

func parseLine(line string) instruction {
	x := strings.Split(line, " ")
	op := x[0]
	arg, _ := strconv.Atoi(x[1])
	return instruction{op, arg}

}

func execute(ins []instruction) (int, bool) {
	acc := 0
	curPos := 0
	visitedPos := NewSet()
	loopDetected := false
	for !loopDetected {
		if visitedPos.Contains(curPos) { // loop detected
			loopDetected = true
			break
		}
		visitedPos.Add(curPos)
		in := ins[curPos]
		switch in.op {
		case "nop":
			curPos++
		case "acc":
			acc += in.arg
			curPos++
		case "jmp":
			curPos += in.arg
		}
		if curPos == len(ins) { //reached EOF
			break
		}
	}
	return acc, loopDetected
}

// Set https://www.davidkaya.com/sets-in-golang/

var exists = struct{}{} // empty struct

type set struct {
	m map[int]struct{}
}

// NewSet returns a et of ints
func NewSet() *set {
	s := &set{}
	s.m = make(map[int]struct{})
	return s
}

func (s *set) Add(value int) {
	s.m[value] = exists
}

func (s *set) Remove(value int) {
	delete(s.m, value)
}

func (s *set) Contains(value int) bool {
	_, c := s.m[value]
	return c
}
