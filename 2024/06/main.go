package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

type direction int

const (
	up direction = iota
	down
	left
	right
)

func (d direction) turn() direction {
	switch d {
	case up:
		return right
	case right:
		return down
	case down:
		return left
	case left:
		return up
	default:
		panic(fmt.Sprintf("unknown direction: %v", d))
	}
}

func (d direction) next(r, c int) (int, int) {
	switch d {
	case up:
		return r - 1, c
	case right:
		return r, c + 1
	case down:
		return r + 1, c
	case left:
		return r, c - 1
	default:
		panic(fmt.Sprintf("unknown direction: %v", d))
	}
}

func step(room roomMap, dir direction, r int, c int, visited map[string]bool) int {
	if room.outside(r, c) {
		return 0
	}

	count := 0
	key := fmt.Sprintf("%d-%d", r, c)

	if !visited[key] {
		visited[key] = true
		count++
	}

	newR, newC := dir.next(r, c)

	if !room.outside(newR, newC) && room[newR][newC] == "#" {
		dir = dir.turn()
		newR, newC = dir.next(r, c)
	}

	return count + step(room, dir, newR, newC, visited)
}

func stepWithLoopCheck(room roomMap, dir direction, r int, c int, visited map[string]bool) int {
	if room.outside(r, c) {
		return 0
	}
	//if there is a wall to the right check if there is a loop if you added a wall to the next spot

}

type roomMap [][]string

func (r roomMap) findGuard() (int, int) {
	idx := -1
	rowNum := -1
	for i, row := range r {
		rowNum = i
		if idx = slices.Index(row, "^"); idx != -1 {
			break
		}
	}
	return rowNum, idx
}

func (r roomMap) hallwayHasWall(dir direction, row int, col int) bool {
	right := dir.turn()
	nr, nc := right.next(row, col)
	for !r.outside(nr, nc) {
		if r[nr][nc] == "#" {
			return true
		}
		nr, nc = right.next(row, col)
	}
	return false
}

func (r roomMap) outside(row, col int) bool {
	return row < 0 || row >= len(r) || col >= len((r)[0]) || col < 0
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()
	s := bufio.NewScanner(file)

	room := roomMap{}
	for s.Scan() {
		room = append(room, strings.Split(s.Text(), ""))
	}

	fmt.Printf("part 1: %d\n", part1(room))

	fmt.Printf("part 2: %d\n", part2(room))
}

func part1(room roomMap) int {
	r, c := room.findGuard()

	return step(room, up, r, c, make(map[string]bool))
}

func part2(room roomMap) int {
	//for each step if there is an object to the right of the guard check for loop
	return -1
}
