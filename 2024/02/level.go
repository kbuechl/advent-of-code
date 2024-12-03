package main

import (
	"strconv"
	"strings"
)

type report []string

func newReport(s string) report {
	return strings.Split(s, " ")
}

func (r report) direction() direction {
	a, b := r.level(0), r.level(1)
	if a == b {
		return equal
	}

	if a > b {
		return decreasing
	}

	return increasing
}

func (r report) isSafe() bool {
	for i := 0; i < len(r)-1; i++ {
		if !r.isLevelSafe(i, i+1) {
			return false
		}
	}

	return true
}

func (r report) level(i int) int {
	l, err := strconv.Atoi(r[i])
	if err != nil {
		panic(err)
	}

	return l
}

func (r report) isLevelSafe(i, j int) bool {
	left, right := r.level(i), r.level(j)

	if abs(left, right) > 3 {
		return false
	}

	switch r.direction() {
	case decreasing:
		return left > right
	case increasing:
		return right > left
	case equal:
		return false //never safe
	default:
		panic("direction not found")
	}
}

func (r report) dampen() bool {
	//find first not safe and remove it
	safe := r
	for i := 0; i < len(r)-1; i++ {
		if !r.isLevelSafe(i, i+1) {
			if len(r) > i+2 && r.isLevelSafe(i, i+2) {
				safe = r.remove(i + 1)
			} else {
				safe = r.remove(i)
			}
			break
		}
	}

	return safe.isSafe()
}

func (r report) remove(idx int) report {
	return append(r[:idx], r[idx+1:]...)
}
