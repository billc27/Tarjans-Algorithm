package main

import (
	"fmt"
	"os"
	"os/exec"
	"image"
	"image/draw"
	"image/png"

	"github.com/dominikbraun/graph"
	graphdraw "github.com/dominikbraun/graph/draw"
)

func drawInputGraph(edges [][2]string) {
	g := graph.New(graph.StringHash, graph.Directed())

	for _, edge := range edges {
		_ = g.AddVertex(edge[0])
		_ = g.AddVertex(edge[1])
		_ = g.AddEdge(edge[0], edge[1])
	}

	file, _ := os.Create("input-graph.gv")
	_ = graphdraw.DOT(g, file)

	cmd := exec.Command("dot", "-Tpng", "-O", "input-graph.gv")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

func drawSCCGraph(scc [][]string) {
	g := graph.New(graph.StringHash, graph.Directed())

	for _, component := range scc {
		for _, vertex := range component {
			_ = g.AddVertex(vertex)
		}
		for i := 0; i < len(component)-1; i++ {
			_ = g.AddEdge(component[i], component[i+1])
		}
		if len(component) > 1 {
			_ = g.AddEdge(component[len(component)-1], component[0])
		}
	}

	file, _ := os.Create("scc-graph.gv")
	_ = graphdraw.DOT(g, file)

	cmd := exec.Command("dot", "-Tpng", "-O", "scc-graph.gv")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

func combineImages(imagePaths []string) (*image.RGBA, error) {
	var images []image.Image
	totalWidth := 0
	maxHeight := 0

	for _, imagePath := range imagePaths {
		file, err := os.Open(imagePath)
		if err != nil {
			return nil, err
		}
		defer file.Close()

		img, _, err := image.Decode(file)
		if err != nil {
			return nil, err
		}

		bounds := img.Bounds()
		totalWidth += bounds.Dx()
		if bounds.Dy() > maxHeight {
			maxHeight = bounds.Dy()
		}

		images = append(images, img)
	}

	result := image.NewRGBA(image.Rect(0, 0, totalWidth, maxHeight))
	currentX := 0

	for _, img := range images {
		bounds := img.Bounds()
		draw.Draw(result, image.Rect(currentX, 0, currentX+bounds.Dx(), bounds.Dy()), img, bounds.Min, draw.Src)
		currentX += bounds.Dx()
	}

	return result, nil
}

func drawBridgeGraph(bridges [][2]string) {
	imagePaths := []string{}

	for i, bridge := range bridges {
		g := graph.New(graph.StringHash, graph.Directed())
		_ = g.AddVertex(bridge[0])
		_ = g.AddVertex(bridge[1])
		_ = g.AddEdge(bridge[0], bridge[1])

		fileName := fmt.Sprintf("bridge-graph-%d.gv", i)
		file, _ := os.Create(fileName)
		_ = graphdraw.DOT(g, file)

		cmd := exec.Command("dot", "-Tpng", "-O", fileName)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		err := cmd.Run()
		if err != nil {
			panic(err)
		}

		imagePaths = append(imagePaths, fmt.Sprintf("bridge-graph-%d.gv.png", i))
	}

	result, err := combineImages(imagePaths)
	if err != nil {
		panic(err)
	}

	file, err := os.Create("bridges-graph.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = png.Encode(file, result)
	if err != nil {
		panic(err)
	}
}

// func main() {
// 	inputGraph := [][2]string{
// 		{"A", "B"},
// 		{"B", "C"},
// 		{"C", "A"},
// 		{"B", "D"},
// 		{"D", "E"},
// 		{"E", "F"},
// 		{"F", "E"},
// 	}

// 	scc := [][]string{
// 		{"F", "E"},
// 		{"D"},
// 		{"C", "B", "A"},
// 	}

// 	bridge := [][2]string{
// 		{"E", "F"},
// 		{"D", "E"},
// 		{"B", "D"},
// 	}

// 	drawInputGraph(inputGraph)
// 	drawSCCGraph(scc)
// 	drawBridgeGraph(bridge)
// }