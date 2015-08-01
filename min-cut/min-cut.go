package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type Edge struct {
	x, y int
}

type Graph struct {
	nodes int
	edges []Edge
}

func print(graph *Graph) {
	fmt.Println("Nodes : ", graph.nodes)
	for i := range graph.edges {
		fmt.Println(graph.edges[i])
	}
	fmt.Println("--------")
}

func minVertex(x Edge) int {
	if x.x < x.y {
		return x.x
	} else {
		return x.y
	}
}

func reduce(graph *Graph) *Graph {
	nodes := graph.nodes - 1

	mergeEdge := graph.edges[rand.Intn(len(graph.edges))]

	edges := make([]Edge, 0, 1000)
	minCord := minVertex(mergeEdge)
	for _, edge := range graph.edges {

		x, y := edge.x, edge.y
		if x == mergeEdge.x || x == mergeEdge.y {
			x = minCord
		}

		if y == mergeEdge.x || y == mergeEdge.y {
			y = minCord
		}

		if x != y {
			edges = append(edges, Edge{x, y})
		}

	}

	return &Graph{nodes: nodes, edges: edges}
}

func reduceSolve(graph *Graph) int {
	for graph.nodes > 2 {
		graph = reduce(graph)
	}
	return len(graph.edges)
}

func solve(graph *Graph) {
	min := 200
	cuts := 0
	for i := 0; i < 200; i++ {
		cuts = reduceSolve(graph)
		if min > cuts {
			min = cuts
		}
	}
	fmt.Println("Cuts : ", min)
}

func main() {
	n := 200

	edges := make([]Edge, 0, 1000)
	in := bufio.NewReader(os.Stdin)
	for i := 0; i < n; i++ {
		line, _ := in.ReadString('\n')
		line = strings.Trim(line, "\t\n\r")

		splits := strings.Split(line, "\t")
		left, _ := strconv.Atoi(splits[0])
		for j := 1; j < len(splits); j++ {
			val, _ := strconv.Atoi(splits[j])
			if left != val {
				edge := Edge{left - 1, val - 1}
				edges = append(edges, edge)
			}
		}
	}

	solve(&Graph{n, edges})
}
