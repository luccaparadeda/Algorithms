package main

import (
	"fmt"
	"strconv"

	graph "github.com/luccaparadeda/Algorithms/Graph"
)

type DFS struct {
	visited map[*graph.Node]bool
	found   bool
}

func main() {
	g := graph.NewGraph()

	fmt.Println(g)

	dfs := &DFS{visited: make(map[*graph.Node]bool)}
	path := ""
	dfs.DFSRecursion(g, g.Nodes[0], 2, nil, &path)

	fmt.Println(path)
}

func (dfs *DFS) DFSRecursion(graph *graph.Graph, node *graph.Node, target int64, parent *graph.Node, path *string) {
	if dfs.visited[node] || dfs.found {
		return
	}

	if *path != "" {
		*path += " -> "
	}

	nodeValue := strconv.FormatInt(node.Value, 10)
	if node.Value == target {
		// Turn the text green
		nodeValue = "\033[32m" + nodeValue + "\033[0m"
	}

	*path += nodeValue

	if node.Value == target {
		dfs.found = true
		return
	}

	dfs.visited[node] = true
	for _, neighbor := range node.Points {
		if !dfs.visited[neighbor] {
			dfs.DFSRecursion(graph, neighbor, target, node, path)
		}
	}
	if parent != nil {
		dfs.DFSRecursion(graph, parent, target, nil, path)
	}
}
