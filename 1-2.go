package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.ReadFile("1.txt")
	if err != nil {
		os.Exit(-1)
	}

	answer := 0
	current := 50
	for i := range strings.SplitSeq(string(f[:len(f)-1]), "\n") {
		direction := 1
		if i[0] == 'L' {
			direction = -1
		}

		amount, _ := strconv.Atoi(i[1:])
		current += amount * direction

		switch {
		case current < 0:
			if current-direction*amount != 0 {
				answer++
			}

			answer += (-current) / 100
			current = (100 - (-current)%100) % 100
		case current > 99:
			answer += current / 100
			current %= 100
		case current == 0:
			answer++
		}
	}

	fmt.Println(answer)
}
