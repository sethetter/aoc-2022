package main

import (
	"bytes"
	"fmt"
	"log"

	h "github.com/sethetter/aoc-2022/helpers"
)

func main() {
	input := h.GetInput("day04")
	fmt.Printf("part 1: %d", part1(input))
}

// For each pair, how many fully contain the other?
func part1(input []byte) int {
	lines := bytes.Split(bytes.TrimSpace(input), []byte("\n"))
	count := 0
	for _, line := range lines {
		s1, s2 := parseLine(line)
		if aContainsB(s1, s2) || aContainsB(s2, s1) {
			count += 1
		}
	}
	return count
}

type section [2]int

func parseLine(line []byte) (section, section) {
	parts := bytes.Split(line, []byte(","))
	if len(parts) != 2 {
		log.Fatalf("expected exactly two sections in line: %s", line)
	}
	return parseSection(parts[0]), parseSection(parts[1])
}

func parseSection(in []byte) section {
	parts := bytes.Split(in, []byte("-"))
	if len(parts) != 2 {
		log.Fatalf("expected exactly two parts of a range: %s", in)
	}
	start, end := h.BytesToInt(parts[0]), h.BytesToInt(parts[1])
	return section{start, end}
}

func aContainsB(a, b section) bool {
	return a[0] <= b[0] && a[1] >= b[1]
}
