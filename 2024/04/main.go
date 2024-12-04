package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type searchDirection int

const (
	right searchDirection = iota
	left
	up
	down
	diagonalBottomRight
	diagonalTopLeft
	diagonalBottomLeft
	diagonalTopRight
)

// returns row and column modifier
func (d searchDirection) coords() (int, int) {
	switch d {
	case up:
		return -1, 0
	case down:
		return 1, 0
	case right:
		return 0, 1
	case left:
		return 0, -1
	case diagonalBottomRight:
		return 1, 1
	case diagonalTopLeft:
		return -1, -1
	case diagonalBottomLeft:
		return 1, -1
	case diagonalTopRight:
		return -1, 1
	default:
		panic("invalid searchDfs direction")
	}
}

func (d searchDirection) mirror() searchDirection {
	switch d {
	case diagonalBottomRight:
		return diagonalBottomLeft
	case diagonalBottomLeft:
		return diagonalBottomRight
	case diagonalTopLeft:
		return diagonalTopRight
	case diagonalTopRight:
		return diagonalTopLeft
	default:
		panic("only diagonal supported")
	}
}

type matrix [][]string

func (m matrix) within(r, c int) bool {
	if r < 0 || r >= len(m) || c < 0 || c >= len(m[0]) {
		return false
	}
	return true
}

func (m matrix) searchDfs(direction searchDirection, target string, r int, c int, idx int) bool {
	if idx >= len(target) {
		return true
	}

	t := string(target[idx])
	x, y := direction.coords()

	nrow, ncol := r+x, c+y
	if !m.within(nrow, ncol) {
		return false
	}

	if m[nrow][ncol] == t {
		return m.searchDfs(direction, target, nrow, ncol, idx+1)
	}
	return false
}

func (m matrix) count(target string, r int, c int) int {
	if m[r][c] != string(target[0]) {
		return 0
	}

	count := 0
	dir := []searchDirection{up, down, left, right, diagonalBottomRight, diagonalBottomLeft, diagonalTopRight, diagonalTopLeft}
	for _, d := range dir {
		if m.searchDfs(d, target, r, c, 1) {
			count++
		}
	}
	return count
}

func part1(m matrix) int {
	const key = "XMAS"

	count := 0

	for row := 0; row < len(m); row++ {
		for col := 0; col < len(m[0]); col++ {
			count += m.count(key, row, col)
		}
	}
	return count
}

func part2(m [][]string) int {
	SAM := "SAM"
	MAS := "MAS"

	numCols := len(m[0])
	numRows := len(m)

	count := 0
	var dfs func(direction searchDirection, key string, r int, c int, idx int) bool

	dfs = func(direction searchDirection, key string, r int, c int, idx int) bool {
		if idx >= len(key) {
			return true
		}

		target := string(key[idx])

		x, y := direction.coords()
		nrow, ncol := r+x, c+y
		if nrow < 0 || nrow >= numRows || ncol < 0 || ncol >= numCols {
			return false
		}
		if m[nrow][ncol] == target {
			return dfs(direction, key, nrow, ncol, idx+1)
		}

		return false
	}

	d := diagonalBottomRight
	//we dont need to searchDfs last 3 rows or last 2 columns
	for row := 0; row < len(m); row++ {
		for col := 0; col < len(m[0]); col++ {
			//find x and searchDfs for next
			c := col + 2

			if m[row][col] == string(SAM[0]) {
				if dfs(d, SAM, row, col, 1) && (m[row][c] == "S" && (dfs(diagonalBottomLeft, SAM, row, c, 1)) || (m[row][c] == "M" && dfs(diagonalBottomLeft, MAS, row, c, 1))) {
					count++
				}
			} else if m[row][col] == string(MAS[0]) {
				if dfs(d, MAS, row, col, 1) && (m[row][c] == "S" && (dfs(diagonalBottomLeft, SAM, row, c, 1)) || (m[row][c] == "M" && dfs(diagonalBottomLeft, MAS, row, c, 1))) {
					count++
				}
			}
		}
	}
	return count
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()
	s := bufio.NewScanner(file)

	m := make([][]string, 0)
	for s.Scan() {
		m = append(m, strings.Split(s.Text(), ""))
	}

	fmt.Printf("part 1: %d\n", part1(m))
	fmt.Printf("part 2: %d\n", part2(m))
}
