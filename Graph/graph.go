package graph

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Value  int64
	Points []*Node
}

type Graph struct {
	Nodes   []*Node
	nodeMap map[int64]*Node
}

func (graph *Graph) String() string {
	str := ""
	for _, node := range graph.Nodes {
		str += strconv.FormatInt(node.Value, 10) + " -> "
		for _, point := range node.Points {
			str += strconv.FormatInt(point.Value, 10) + " "
		}
		str += "\n"
	}
	return str
}

func (graph *Graph) getOrCreateNode(value int64) *Node {
	if node, exists := graph.nodeMap[value]; exists {
		return node
	}
	node := &Node{Value: value}
	graph.nodeMap[value] = node
	graph.Nodes = append(graph.Nodes, node)
	return node
}

func NewGraph() *Graph {
	file, err := os.Open("./graph.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	graph := &Graph{nodeMap: make(map[int64]*Node)}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		nodes := strings.Split(line, " ")

		nodeValue, err := strconv.ParseInt(nodes[0], 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		node := graph.getOrCreateNode(nodeValue)

		for _, point := range nodes[1:] {
			pointValue, err := strconv.ParseInt(point, 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			neighbor := graph.getOrCreateNode(pointValue)
			node.Points = append(node.Points, neighbor)
		}
	}

	return graph
}
