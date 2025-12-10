// This solution was either stolen, adapted or translated from another language by me.
// I keep track of tasks which I failed to solve and this is one of them.
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func applyButtons(diagram []bool, buttons []int) []bool {
	out := make([]bool, len(diagram))
	copy(out, diagram)

	for _, i := range buttons {
		if i >= 0 && i < len(out) {
			out[i] = !out[i]
		}
	}

	return out
}

func sameDiagrams(a, b []bool) bool {
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func diagramToString(diagram []bool) string {
	buffer := make([]byte, len(diagram))

	for i := range diagram {
		if diagram[i] {
			buffer[i] = '1'
			continue
		}

		buffer[i] = '0'
	}

	return string(buffer)
}

func calculateMinPresses(target []bool, buttons [][]int) int {
	type node struct {
		diagrams []bool
		steps    int
	}

	startState := make([]bool, len(target))
	queue := []node{{startState, 0}}
	seen := map[string]bool{diagramToString(startState): true}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if sameDiagrams(current.diagrams, target) {
			return current.steps
		}

		for _, button := range buttons {
			next := applyButtons(current.diagrams, button)
			k := diagramToString(next)

			if !seen[k] {
				seen[k] = true
				queue = append(queue, node{next, current.steps + 1})
			}
		}
	}

	return -1
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		os.Exit(-1)
	}

	total := 0
	for line := range strings.SplitSeq(strings.TrimSpace(string(data)), "\n") {
		parts := strings.Fields(line)

		rawDiagram := parts[0][1 : len(parts[0])-1]
		diagram := make([]bool, len(rawDiagram))
		for i, c := range rawDiagram {
			diagram[i] = c == '#'
		}

		buttons := [][]int{}
		for _, parentheses := range parts[1:] {
			if strings.HasPrefix(parentheses, "{") {
				break
			}

			parentheses = strings.Trim(parentheses, "()")
			numbers := strings.Split(parentheses, ",")
			button := make([]int, 0, len(numbers))

			for _, ns := range numbers {
				if ns == "" {
					continue
				}

				n, _ := strconv.Atoi(ns)
				button = append(button, n)
			}

			buttons = append(buttons, button)
		}

		total += calculateMinPresses(diagram, buttons)
	}

	fmt.Println(total)
}
