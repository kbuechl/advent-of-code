package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()
	s := bufio.NewScanner(file)
	fmt.Println(part1(s))

	//reset file from earlier search
	if _, err = file.Seek(0, 0); err != nil {
		panic(err)
	}

	s2 := bufio.NewScanner(file)
	fmt.Println(part2(s2))
}

func part1(scanner *bufio.Scanner) string {
	panic("not implemented")
}

func part2(scanner *bufio.Scanner) string {
	panic("not implemented")
}
