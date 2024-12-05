package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type set map[string]bool

func newSet(keys []string) set {
	s := make(set)
	for _, k := range keys {
		s.add(k)
	}
	return s
}

func (s set) intersect(s2 set) set {
	intersect := set{}
	for i, _ := range s {
		if s2[i] {
			intersect[i] = true
		}
	}
	return intersect
}

func (s set) remove(k string) {
	delete(s, k)
}

func (s set) add(k string) {
	s[k] = true
}

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

func part1(scanner *bufio.Scanner) int {
	//read until the empty line
	//create a map of pages and a set of the pages they are dependent on
	validUpdates := make([][]string, 0)

	m := make(map[string]set)
	for scanner.Scan() {
		l := scanner.Text()
		if l == "" {
			break
		}
		split := strings.Split(l, "|")
		l, r := split[0], split[1] //l is the page, r is the requirement
		if _, ok := m[l]; !ok {
			m[l] = set{}
		}
		m[l].add(r)
	}

	//we should be after the new line now, so each line is an update
	for scanner.Scan() {
		l := scanner.Text()
		split := strings.Split(l, ",")
		if isInOrder(split, m) {
			validUpdates = append(validUpdates, split)
		}
	}

	sum := 0
	for _, u := range validUpdates {
		page, _ := strconv.Atoi(u[len(u)/2])
		sum += page
	}
	return sum
}

func isInOrder(pageUpdate []string, m map[string]set) bool {
	s := newSet(pageUpdate)
	for i := len(pageUpdate) - 1; i >= 0; i-- {
		cur := pageUpdate[i]
		if s2, ok := m[cur]; ok {
			if len(s.intersect(s2)) > 0 {
				return false
			}
		}
		s.remove(cur)
	}
	return true
}

func part2(scanner *bufio.Scanner) int {
	return -1
}
