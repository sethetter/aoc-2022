package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_startOfPacketIdx(t *testing.T) {
	cases := []struct {
		in     []byte
		expect int
	}{
		{[]byte("bvwbjplbgvbhsrlpgdmjqwftvncz"), 5},
		{[]byte("nppdvjthqldpwncqszvftbrmjlhg"), 6},
		{[]byte("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"), 10},
		{[]byte("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"), 11},
	}
	for _, tc := range cases {
		assert.Equal(t, tc.expect, startOfPacketIdx(tc.in))
	}
}
