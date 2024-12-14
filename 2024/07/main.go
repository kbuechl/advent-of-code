package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func newEquation(s string) equation {
	e := equation{}
	split := strings.Split(s, ":")
	e.value, _ = strconv.Atoi(split[0])

	nSplit := strings.Split(strings.Trim(split[1]," "), " ")
	for _, str := range nSplit {
		n, _ := strconv.Atoi(str)
		e.numbers = append(e.numbers, n)
	}
	return e
}

type equation struct {
	value   int
	numbers []int
}

type equationList []equation

func(el equationList) sumValid(operators []string) int {
	testValues := make([]int, 0)

	for _, e := range el {
		count := calculate(e.value, e.numbers[0], e.numbers[1:], operators)
		if count >0 {
			testValues = append(testValues, e.value)
		}
	}

	sum := 0
	for _, v := range testValues {
		sum += v
	}

	return sum
}

func calculate(value int, curTotal int, numbers []int, operators []string) int {
	if len(numbers) == 0 {
		if curTotal == value {
			return 1
		}
		return 0
	}

	applyOperator := func(operator string, currentTotal, nextNumber int) int {
		switch operator {
		case "+":
			return currentTotal + nextNumber
		case "*":
			return currentTotal * nextNumber
		case "||":
			concatenated := fmt.Sprintf("%d%d", currentTotal, nextNumber)
			result, _ := strconv.Atoi(concatenated)
			return result
		default:
			panic("operator not found")
		}
	}

	count := 0
	nextNumbers := numbers[1:]

	for _, operator := range operators {
		result := applyOperator(operator, curTotal, numbers[0])
		count += calculate(value, result, nextNumbers, operators)
	}

	return count
}

func part1(equations equationList) int {
	return equations.sumValid([]string{"*", "+"} )
}

func part2(equations equationList) int {
	return equations.sumValid([]string{"*", "+", "||"}  )
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	defer file.Close()
	s := bufio.NewScanner(file)

	equations := make(equationList, 0)
	for s.Scan() {
		l := s.Text()
		equations = append(equations, newEquation(l))
	}
	fmt.Printf("part 1: %d\n", part1(equations))

	fmt.Printf("part 2: %d\n", part2(equations))
}
