package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type direction int

const (
	increasing direction = iota
	decreasing
	equal
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()
	s := bufio.NewScanner(file)
	fmt.Printf("part 1: %d\n", part1(s))

	//reset file from earlier search
	if _, err = file.Seek(0, 0); err != nil {
		panic(err)
	}

	s2 := bufio.NewScanner(file)
	fmt.Printf("part 2: %d\n", part2(s2))
}

func part1(s *bufio.Scanner) int {
	count := 0
	for s.Scan() {
		r := newReport(s.Text())
		if r.isSafe() {
			count++
		}
	}
	return count
}

func part2(s *bufio.Scanner) int {
	count := 0
	for s.Scan() {
		r := newReport(s.Text())
		if r.dampen() {
			count++
		}
	}
	return count
}

func abs(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

func getDirection(slice []string) direction {
	a, err := strconv.Atoi(slice[0])
	if err != nil {
		panic(err)
	}
	b, err := strconv.Atoi(slice[1])
	if err != nil {
		panic(err)
	}

	if a > b {
		return decreasing
	}
	return increasing
}

func isSafe(l int, r int, dir direction) bool {
	if (dir == increasing && r < l) || (dir == decreasing && l < r) {
		return false
	}
	if abs(l, r) == 0 || abs(l, r) > 3 {
		return false
	}
	return true
}
