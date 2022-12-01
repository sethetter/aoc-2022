package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("day01/input.txt")
	if err != nil {
		log.Fatal("failed to read input file")
	}

	counts := calorieCounts(string(input))
	most := mostCalories(counts)
	fmt.Printf("part 1: %d\n", most)
}

func calorieCounts(input string) []int {
	itemLists := strings.Split(
		strings.TrimSpace(string(input)),
		"\n\n",
	)

	counts := make([]int, len(itemLists))

	for i, list := range itemLists {
		items := strings.Split(list, "\n")
		for _, item := range items {
			x, err := strconv.Atoi(item)
			if err != nil {
				log.Fatalf("failed to parse as int: %s", item)
			}
			counts[i] += x
		}
	}

	return counts
}

func mostCalories(counts []int) int {
	m := 0
	for _, c := range counts {
		if c > m {
			m = c
		}
	}
	return m
}
