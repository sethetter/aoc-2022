package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("day02/input.txt")
	if err != nil {
		log.Fatal("failed to read input file")
	}

	fmt.Println(calculateScore(string(input), win_p2))
}

type play int

const (
	rock play = iota
	paper
	scissors
)

func letterToPlay(l string) play {
	switch strings.ToUpper(strings.TrimSpace(l)) {
	case "A":
		return rock
	case "X":
		return rock
	case "B":
		return paper
	case "Y":
		return paper
	case "C":
		return scissors
	case "Z":
		return scissors
	}
	panic(fmt.Sprintf("unexpected play: %s", l))
}

type result int

const (
	win_p1 result = iota
	win_p2
	draw
)

func whoWon(p1, p2 play) result {
	if p1 == p2 {
		return draw
	}

	switch p1 {
	case rock:
		switch p2 {
		case paper:
			return win_p2
		case scissors:
			return win_p1
		}
	case paper:
		switch p2 {
		case rock:
			return win_p1
		case scissors:
			return win_p2
		}
	case scissors:
		switch p2 {
		case rock:
			return win_p2
		case paper:
			return win_p1
		}
	}
	panic("unreachable")
}

func scoreForPlay(p play) int {
	switch p {
	case rock:
		return 1
	case paper:
		return 2
	case scissors:
		return 3
	}
	panic("unexpected play")
}

func calculateScore(strategy string, player result) int {
	rounds := strings.Split(strings.TrimSpace(strategy), "\n")

	score := 0
	for _, round := range rounds {
		plays := strings.Split(round, " ")
		p1, p2 := letterToPlay(plays[0]), letterToPlay(plays[1])

		switch player {
		case win_p1:
			score += scoreForPlay(p1)
		case win_p2:
			score += scoreForPlay(p2)
		}

		switch whoWon(p1, p2) {
		case player:
			score += 6
		case draw:
			score += 3
		}
	}

	return score
}
