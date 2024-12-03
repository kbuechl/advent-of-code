package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	fmt.Printf("part 2:%d\n", part2(s2))
}

func part1(scanner *bufio.Scanner) int {
	lHeap := &minHeap{}
	rHeap := &minHeap{}

	heap.Init(lHeap)
	heap.Init(rHeap)

	//parse input & build heaps
	for scanner.Scan() {
		line := scanner.Text()
		l, r := parse(line)
		heap.Push(lHeap, l)
		heap.Push(rHeap, r)
	}

	if lHeap.Len() != rHeap.Len() {
		panic("heap size differ")
	}

	dist := 0
	//calc distance
	for lHeap.Len() > 0 {
		l, r := heap.Pop(lHeap).(int), heap.Pop(rHeap).(int)
		dist += distance(l, r)
	}

	return dist
}

func part2(scanner *bufio.Scanner) int {
	count := make(map[int]int)
	items := make([]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		l, r := parse(line)
		items = append(items, l)
		count[r]++
	}

	score := 0
	for _, n := range items {
		score += n * count[n]
	}

	return score
}

func distance(i1, i2 int) int {
	if i1 >= i2 {
		return i1 - i2
	}
	return i2 - i1
}

func parse(line string) (int, int) {
	data := strings.Split(line, "   ")
	if len(data) != 2 {
		panic("invalid input")
	}
	l, err := strconv.Atoi(data[0])
	if err != nil {
		panic(err)
	}
	r, err := strconv.Atoi(data[1])
	if err != nil {
		panic(err)
	}
	return l, r
}
