package main

import (
	"bufio"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestSolution(t *testing.T) {
	file, err := os.Open("example.txt")
	require.NoError(t, err)
	defer file.Close()
	s := bufio.NewScanner(file)

	t.Run("part one", func(t *testing.T) {
		assert.Equal(t, 3749, part1(s))
	})

	t.Run("part two", func(t *testing.T) {
		assert.Equal(t, 11387, part2(s))
	})
}
