package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

const n int = 200

type Item struct {
	node  int
	value int
}

type Edge struct {
	x, y   int
	weight int
}

type PriorityQueue []Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].value < pq[j].value
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := pq.Len()
	item := x.(Item)
	*pq = append(*pq, item)
	for (*pq)[n].value < (*pq)[n/2].value && n >= 1 {
		(*pq)[n], (*pq)[n/2] = (*pq)[n/2], (*pq)[n]
		n = n / 2
	}
}

func (pq *PriorityQueue) Min() interface{} {
	n := pq.Len()
	var ret Item = (*pq)[1]
	(*pq)[1], (*pq)[n-1] = (*pq)[n-1], (*pq)[1]
	*pq = (*pq)[0 : n-1]
	i := 1
	for {
		var a, b int = math.MaxInt32, math.MaxInt32
		var ind int

		if i*2 < n-1 {
			a = (*pq)[i*2].value
		}
		if i*2+1 < n-1 {
			b = (*pq)[i*2+1].value
		}

		if a < b {
			ind = i * 2
		} else {
			ind = i*2 + 1
		}

		if a == math.MaxInt32 && b == math.MaxInt32 {
			break
		}

		if (*pq)[ind].value < (*pq)[i].value {
			(*pq)[ind], (*pq)[i] = (*pq)[i], (*pq)[ind]
			i = ind
		} else {
			break
		}
	}
	return ret
}

func solve(n int, adjList [][]Edge) {
	s := 0
	p := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = math.MaxInt32
	}

	var pq PriorityQueue = make([]Item, 1, 1000)

	pq.Push(Item{s, 0})

	for len(pq) > 1 {
		node := pq.Min().(Item)

		if p[node.node] != math.MaxInt32 {
			continue
		}
		p[node.node] = node.value
		for i := 0; i < len(adjList[node.node]); i++ {
			edge := adjList[node.node][i]
			if p[edge.y] == math.MaxInt32 {
				item := Item{edge.y, edge.weight + node.value}
				pq.Push(item)
			}
		}

	}

	arr := []int{7, 37, 59, 82, 99, 115, 133, 165, 188, 197}
	for j := 0; j < len(arr); j++ {
		fmt.Printf("%d,", p[arr[j]-1])
	}

}

func main() {
	n := 200

	adjList := make([][]Edge, n)

	for i := 0; i < n; i++ {
		adjList[i] = make([]Edge, 0, 200)
	}

	in := bufio.NewReader(os.Stdin)
	for {
		line, err := in.ReadString('\n')
		if err == io.EOF {
			break
		}

		line = strings.Trim(line, "\t\n\r")
		splits := strings.Split(line, "\t")

		u, _ := strconv.Atoi(splits[0])

		for j := 1; j < len(splits); j++ {
			ev := splits[j]
			sp := strings.Split(ev, ",")

			v, _ := strconv.Atoi(sp[0])
			w, _ := strconv.Atoi(sp[1])

			adjList[u-1] = append(adjList[u-1], Edge{u - 1, v - 1, w})
		}
	}

	solve(n, adjList)
}
