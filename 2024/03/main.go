package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var (
	validInput   = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	disableRegex = regexp.MustCompile(`don't\(\).*?do\(\)`)
)

func init() {

}
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

func part1(scanner *bufio.Scanner) int {
	answer := 0
	for scanner.Scan() {
		text := scanner.Text()
		answer += calculateMatches(text)
	}

	return answer
}

func calculateMatches(text string) int {
	res := 0
	matches := validInput.FindAllStringSubmatch(text, -1)
	if len(matches) == 0 {
		panic("no matches found")
	}

	for _, match := range matches {
		if len(match) != 3 {
			panic("match length does not meet criteria")
		}
		x, _ := strconv.Atoi(match[1])
		y, _ := strconv.Atoi(match[2])
		res += x * y
	}
	return res
}

func part2(scanner *bufio.Scanner) int {
	s := ""

	for scanner.Scan() {
		s += scanner.Text()
	}

	s = disableRegex.ReplaceAllString(s, "")
	//todo: get better at regex and catch the edge case
	s = regexp.MustCompile(`don't\(\).*`).ReplaceAllString(s, "")
	return calculateMatches(s)
}
