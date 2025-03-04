// MIT License: See https://github.com/pzsz/voronoi/LICENSE.md

// Author: Przemyslaw Szczepaniak (przeszczep@gmail.com)
// Port of Raymond Hill's (rhill@raymondhill.net) javascript implementation
// of Steven Forune's algorithm to compute Voronoi diagrams

package voronoi_test

import (
	"math/rand"
	"testing"

	. "github.com/haddock7/voronoi"
)

func verifyDiagram(diagram *Diagram, edgesCount, cellsCount, perCellCount int, t *testing.T) {
	if len(diagram.Edges) != edgesCount {
		t.Errorf("Expected %d edges not %d", edgesCount, len(diagram.Edges))
	}

	if len(diagram.Cells) != cellsCount {
		t.Errorf("Expected %d cells not %d", cellsCount, len(diagram.Cells))
	}

	if perCellCount > 0 {
		for _, cell := range diagram.Cells {
			if len(cell.Halfedges) != perCellCount {
				t.Errorf("Expected per cell edge count expected %d, not %d", perCellCount, len(cell.Halfedges))
			}
		}
	}
}

func TestVoronoi2Points(t *testing.T) {
	sites := []SiteVertex{
		{Vertex: Vertex{4, 5}},
		{Vertex: Vertex{6, 5}},
	}

	verifyDiagram(ComputeDiagram(sites, NewBBox(0, 10, 0, 10), true),
		7, 2, 4, t)
	verifyDiagram(ComputeDiagram(sites, NewBBox(0, 10, 0, 10), false),
		1, 2, 1, t)
}

func TestVoronoi3Points(t *testing.T) {
	sites := []SiteVertex{
		{Vertex: Vertex{4, 5}},
		{Vertex: Vertex{6, 5}},
		{Vertex: Vertex{5, 8}},
	}

	verifyDiagram(ComputeDiagram(sites, NewBBox(0, 10, 0, 10), true),
		10, 3, -1, t)
	verifyDiagram(ComputeDiagram(sites, NewBBox(0, 10, 0, 10), false),
		3, 3, 2, t)
}

func Benchmark1000(b *testing.B) {
	rand.Seed(1234567)
	b.StopTimer()
	sites := make([]SiteVertex, 100)
	for j := 0; j < 100; j++ {
		sites[j].Vertex.X = rand.Float64() * 100
		sites[j].Vertex.Y = rand.Float64() * 100
	}
	b.StartTimer()
	ComputeDiagram(sites, NewBBox(0, 100, 0, 100), true)
}
