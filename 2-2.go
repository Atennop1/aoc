package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.ReadFile("2.txt")
	if err != nil {
		os.Exit(1)
	}

	stringDiapasones := strings.Split(string(f[:len(f)-1]), ",")
	diapasones := make([][2]int, len(stringDiapasones))
	for _, i := range stringDiapasones {
		stringBounds := strings.Split(i, "-")
		lowerBounds, _ := strconv.Atoi(stringBounds[0])
		upperBounds, _ := strconv.Atoi(stringBounds[1])
		diapasones = append(diapasones, [2]int{lowerBounds, upperBounds})
	}

	result := 0
	for _, diapasone := range diapasones {
		for i := diapasone[0]; i <= diapasone[1]; i++ {
			iString := strconv.Itoa(i)
			for possibleChunk := 1; possibleChunk <= len(iString)/2; possibleChunk++ {
				if len(iString)%possibleChunk != 0 {
					continue
				}

				if strings.Repeat(iString[:possibleChunk], len(iString)/possibleChunk) == iString {
					result += i
					break
				}
			}
		}
	}

	fmt.Println(result)
}
