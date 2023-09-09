# Voronoi diagrams in Go

A Implementation of of Steven J. Fortune's algorithm to
efficiently compute Voronoi diagrams in Go language. Based on 
a Raymond Hill's javascript implementation 
(https://raw.github.com/gorhill/Javascript-Voronoi).

Forked from github.com/pzsz/voronoi. The only change is a new SiteVertex type that includes
a user data type. Sites were defined by a Vertex struct on the original repository.

## Usage


```go
import "github.com/haddock7/voronoi"

func useVoronoi() {
    // Sites of voronoi diagram
	sites := []SiteVertex{
		{Vertex: Vertex{4, 5}},
		{Vertex: Vertex{6, 5}},
		...
	}

	// Create bounding box of [0, 20] in X axis
	// and [0, 10] in Y axis
	bbox := NewBBox(0, 20, 0, 10)

	// Compute diagram and close cells (add half edges from bounding box)
	diagram := NewVoronoi().ComputeDiagram(sites, bbox, true)

	// Iterate over cells
	for _, cell := diagram.Cells {
		for _, hedge := cell.Halfedges {
		    ...
		}	
	}

	// Iterate over all edges
	for _, edge := diagram.Edge {
	    ...
	}
}