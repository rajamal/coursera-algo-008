package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

const n int = 875714

var time int = 0

func dfs(graph [][]int, node int, visited []bool, list []int) []int {
	visited[node] = true
	for i := 0; i < len(graph[node]); i++ {
		n := graph[node][i]
		if !visited[n] {
			list = dfs(graph, n, visited, list)
		}
	}
	list = append(list, node)
	return list
}

func rank(graph [][]int) []int {
	visited := make([]bool, n)
	for i := 0; i < n; i++ {
		visited[i] = false
	}
	list := make([]int, 0, n)
	for i := n - 1; i >= 0; i-- {
		if !visited[i] {
			list = dfs(graph, i, visited, list)
		}
	}
	return list
}

func dfsForward(graph [][]int, node, source int, visited []bool, count []int) {
	visited[node] = true
	count[source]++
	for i := 0; i < len(graph[node]); i++ {
		n := graph[node][i]
		if !visited[n] {
			dfsForward(graph, n, source, visited, count)
		}
	}
}

func solve(graph, reverseGraph [][]int) {

	list := rank(reverseGraph)

	count := make([]int, n)
	visited := make([]bool, n)

	for i := 0; i < n; i++ {
		count[i] = 0
	}

	for i := n - 1; i >= 0; i-- {
		node := list[i]
		if !visited[node] {
			dfsForward(graph, node, node, visited, count)
		}
	}
	sort.Ints(count)
	for i := 0; i < 5; i++ {
		fmt.Println(count[n-1-i])
	}
}

func main() {

	adjList := make([][]int, n)
	reverseAdjList := make([][]int, n)

	for i := 0; i < n; i++ {
		adjList[i] = make([]int, 0, 10)
		reverseAdjList[i] = make([]int, 0, 10)
	}

	in := bufio.NewReader(os.Stdin)
	for {
		line, err := in.ReadString('\n')
		if err == io.EOF {
			break
		}

		line = strings.Trim(line, "\t\n\r")
		splits := strings.Split(line, " ")

		u, _ := strconv.Atoi(splits[0])
		v, _ := strconv.Atoi(splits[1])

		adjList[u-1] = append(adjList[u-1], v-1)
		reverseAdjList[v-1] = append(reverseAdjList[v-1], u-1)
	}

	solve(adjList, reverseAdjList)
}
