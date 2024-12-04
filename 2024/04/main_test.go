package main

import (
	"bufio"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"strings"
	"testing"
)

func TestSolution(t *testing.T) {
	file, err := os.Open("example.txt")
	require.NoError(t, err)
	defer file.Close()
	s := bufio.NewScanner(file)

	m := make([][]string, 0)
	for s.Scan() {
		m = append(m, strings.Split(s.Text(), ""))
	}
	t.Run("part one", func(t *testing.T) {
		assert.Equal(t, 18, part1(m))
	})

	t.Run("part two", func(t *testing.T) {
		assert.Equal(t, 9, part2(m))
	})
}
