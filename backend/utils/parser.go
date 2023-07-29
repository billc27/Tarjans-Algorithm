package utils

import (
	"strings"
)

func ParseGraph(input string) [][2]string {
	lines := strings.Split(input, "\n")
	inputGraph := make([][2]string, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, " ")
		inputGraph[i] = [2]string{parts[0], parts[1]}
	}
	return inputGraph
}

func ParseSCC(input string) [][]string {
	lines := strings.Split(input, "\n")
	scc := make([][]string, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, " ")
		scc[i] = parts
	}
	return scc
}
