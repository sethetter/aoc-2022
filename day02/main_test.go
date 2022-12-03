package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testInput = `
A Y
B X
C Z
`

func Test_calculateScore(t *testing.T) {
	got := calculateScore(testInput, win_p1)
	assert.Equal(t, 15, got)
}
