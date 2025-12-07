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

	result := 0
	lines = lines[1:]
	beams := map[cell]struct{}{firstBeam: struct{}{}}

	for i, _ := range lines[:len(lines)-1] {
		newBeams := make(map[cell]struct{})
		for k, _ := range beams {
			if lines[i+1][k.x] != '^' {
				newBeams[cell{x: k.x, y: k.y + 1}] = struct{}{}
				continue
			}

			result++
			newBeams[cell{x: k.x - 1, y: k.y + 1}] = struct{}{}
			newBeams[cell{x: k.x + 1, y: k.y + 1}] = struct{}{}
			delete(newBeams, k)
		}

		beams = newBeams
	}

	fmt.Println(result)
}
