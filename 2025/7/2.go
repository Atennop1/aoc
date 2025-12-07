package main

import (
	"fmt"
	"os"
	"strings"
)

type cell struct {
	x int
	y int
}

func addOrSumIfExists(beams map[cell]int, cell cell, value int) {
	if _, exists := beams[cell]; exists {
		beams[cell] += value
		return
	}

	beams[cell] = value
}

func main() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		os.Exit(-1)
	}

	firstBeam := cell{}
	lines := strings.Split(strings.Trim(string(f), "\n"), "\n")

	for i, r := range lines[0] {
		if r == 'S' {
			firstBeam = cell{x: i, y: 0}
			break
		}
	}

	lines = lines[1:]
	beams := map[cell]int{firstBeam: 1}

	for i, _ := range lines[:len(lines)-1] {
		newBeams := make(map[cell]int)
		for k, v := range beams {
			if lines[i+1][k.x] != '^' {
				addOrSumIfExists(newBeams, cell{x: k.x, y: k.y + 1}, v)
				continue
			}

			addOrSumIfExists(newBeams, cell{x: k.x - 1, y: k.y + 1}, v)
			addOrSumIfExists(newBeams, cell{x: k.x + 1, y: k.y + 1}, v)
			delete(newBeams, k)
		}

		beams = newBeams
	}

	result := 0
	for _, v := range beams {
		result += v
	}

	fmt.Println(result)
}
