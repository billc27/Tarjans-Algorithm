package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Graph struct {
	vertices   []string
	edges      map[string][]string
	index      map[string]int
	lowlink    map[string]int
	onStack    map[string]bool
	stack      []string
	indexCount int
	scc        [][]string
}

func NewGraph() *Graph {
	return &Graph{
		vertices:   []string{},
		edges:      make(map[string][]string),
		index:      make(map[string]int),
		lowlink:    make(map[string]int),
		onStack:    make(map[string]bool),
		stack:      []string{},
		indexCount: 0,
		scc:        [][]string{},
	}
}

func (g *Graph) AddEdge(u, v string) {
	if _, ok := g.edges[u]; !ok {
		g.edges[u] = []string{}
		g.vertices = append(g.vertices, u)
	}
	g.edges[u] = append(g.edges[u], v)
}

func (g *Graph) StrongConnect(v string) {
	g.index[v] = g.indexCount
	g.lowlink[v] = g.indexCount
	g.indexCount++
	g.stack = append(g.stack, v)
	g.onStack[v] = true

	for _, w := range g.edges[v] {
		if _, ok := g.index[w]; !ok {
			g.StrongConnect(w)
			if g.lowlink[w] < g.lowlink[v] {
				g.lowlink[v] = g.lowlink[w]
			}
		} else if g.onStack[w] {
			if g.index[w] < g.lowlink[v] {
				g.lowlink[v] = g.index[w]
			}
		}
	}

	if g.lowlink[v] == g.index[v] {
		component := []string{}
		w := ""
		for w != v {
			w, g.stack = g.stack[len(g.stack)-1], g.stack[:len(g.stack)-1]
			g.onStack[w] = false
			component = append(component, w)
		}
		g.scc = append(g.scc, component)
	}
}

func (g *Graph) findSCC() [][]string {
	for _, v := range g.vertices {
		if _, ok := g.index[v]; !ok {
			g.StrongConnect(v)
		}
	}
	return g.scc
}

// func main() {
// 	scanner := bufio.NewScanner(os.Stdin)
// 	graph := NewGraph()
// 	for scanner.Scan() {
//         line := scanner.Text()
//         if line == ""{
//             break
//         }
//         parts := strings.Split(line, " ")
//         graph.AddEdge(parts[0], parts[1])
//     }
// 	scc := graph.findSCC()
// 	for _, component := range scc {
//         fmt.Println(strings.Join(component, " "))
//     }
// }
