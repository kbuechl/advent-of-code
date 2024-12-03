package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var validInput *regexp.Regexp

func init() {
	validInput = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
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
		x, err := strconv.Atoi(match[1])
		if err != nil {
			panic(err)
		}
		y, err := strconv.Atoi(match[2])
		if err != nil {
			panic(err)
		}
		res += x * y
	}
	return res
}
func part2(scanner *bufio.Scanner) int {
	re := regexp.MustCompile(`don't\([^\)]*\)(.*?do\([^\)]*\))?`)
	ans := 0
	for scanner.Scan() {
		text := scanner.Text()
		text = re.ReplaceAllString(text, "")

		ans += calculateMatches(text)
	}

	return ans
}

//func oldpart2(scanner *bufio.Scanner) int {
//	//anything between dont and do can be thrown out
//
//	reDont := regexp.MustCompile(`don't\(\)`)
//	reDo := regexp.MustCompile(`do\(\)`)
//
//	for scanner.Scan() {
//
//	}
//	text := scanner.Text()
//	for {
//		idx := reDont.FindStringIndex(text)
//		id2 := reDo.FindStringIndex(text)
//		if idx == nil {
//
//			break
//		}
//		start := idx[0]
//		stop := len(text) - 1
//		if id2 != nil {
//			stop = id2[1]
//		}
//		text = text[:start-1] + text[stop+1:]
//	}
//
//	return calculateMatches(text)
//}
