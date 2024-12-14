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
	r := roomMap{}

	for s.Scan() {
		r = append(r, strings.Split(s.Text(), ""))
	}

	t.Run("part one", func(t *testing.T) {
		assert.Equal(t, 41, part1(r))
	})

	t.Run("part two", func(t *testing.T) {
		assert.Equal(t, 0, part2(r))
	})
}
