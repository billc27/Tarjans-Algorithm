package utils

import (
	// "bufio"
	// "fmt"
	// "os"
	"strings"
)

type Graph struct {
	vertices []string
	edges    map[string][]string
	time     int
	low      map[string]int
	disc     map[string]int
	parent   map[string]string
	bridges  [][]string
}

func NewUndirectedGraph() *Graph {
	return &Graph{
		vertices: []string{},
		edges:    make(map[string][]string),
		time:     0,
		low:      make(map[string]int),
		disc:     make(map[string]int),
		parent:   make(map[string]string),
		bridges:  [][]string{},
	}
}

func (g *Graph) AddEdge(u, v string) {
	if _, ok := g.edges[u]; !ok {
		g.edges[u] = []string{}
		g.vertices = append(g.vertices, u)
	}
	if _, ok := g.edges[v]; !ok {
		g.edges[v] = []string{}
		g.vertices = append(g.vertices, v)
	}
	g.edges[u] = append(g.edges[u], v)
	g.edges[v] = append(g.edges[v], u)
}

func (g *Graph) BridgeUtil(u string) {
	g.disc[u] = g.time
	g.low[u] = g.time
	g.time++

	for _, v := range g.edges[u] {
		if _, ok := g.disc[v]; !ok {
			g.parent[v] = u
			g.BridgeUtil(v)

			if g.low[v] > g.disc[u] {
				g.bridges = append(g.bridges, []string{u, v})
			}

			if g.low[v] < g.low[u] {
				g.low[u] = g.low[v]
			}
		} else if v != g.parent[u] && g.disc[v] < g.low[u] {
			g.low[u] = g.disc[v]
		}
	}
}

func (g *Graph) findBridge() [][]string {
	for _, v := range g.vertices {
		if _, ok := g.disc[v]; !ok {
			g.BridgeUtil(v)
		}
	}
	return g.bridges
}

// func main() {
	// scanner := bufio.NewScanner(os.Stdin)
	// graph := NewUndirectedGraph()
	// for scanner.Scan() {
    //     line := scanner.Text()
    //     if line == ""{
    //         break
    //     }
    //     parts := strings.Split(line, " ")
    //     graph.AddEdge(parts[0], parts[1])
    // }
	// bridges := graph.findBridge()
	// for _, bridge := range bridges {
	// 	fmt.Println(strings.Join(bridge, " "))
	// }
// }

func GenerateBridges(input string) string {
	graph := NewUndirectedGraph()
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		if line == "" {
			break
		}
		parts := strings.Split(line, " ")
		graph.AddEdge(parts[0], parts[1])
	}
	bridges := graph.findBridge()
	var result []string
	for _, bridge := range bridges {
		result = append(result, strings.Join(bridge, " "))
	}
	return strings.Join(result, "\n")
}