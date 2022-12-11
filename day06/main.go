package main

import (
	"bytes"
	"fmt"

	"github.com/samber/lo"
	h "github.com/sethetter/aoc-2022/helpers"
)

func main() {
	input := bytes.TrimSpace(h.GetInput("day06"))
	fmt.Println(startOfPacketIdx(input))
}

func startOfPacketIdx(in []byte) int {
	for i := range in {
		if len(in) < i+2 || i < 3 {
			continue
		}

		countUniq := 0
		for j := 0; j < 4; j += 1 {
			if len(lo.Uniq([]byte{in[i+j], in[i+j-1], in[i+j-2], in[i+j-3]})) == 4 {
				countUniq += 1
			}
		}
		if countUniq == 4 {
			return i + 1
		}
	}
	return -1
}
