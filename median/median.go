package main

import (
	"fmt"
	"math"
)

type PriorityQueue []int

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i] < pq[j]
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x int) {
	n := pq.Len()
	item := x
	*pq = append(*pq, item)
	for (*pq)[n] < (*pq)[n/2] && n >= 2 {
		(*pq)[n], (*pq)[n/2] = (*pq)[n/2], (*pq)[n]
		n = n / 2
	}
}

func (pq *PriorityQueue) Top() int {
	return (*pq)[1]
}

func (pq *PriorityQueue) Min() int {
	n := pq.Len()
	var ret int = (*pq)[1]
	(*pq)[1], (*pq)[n-1] = (*pq)[n-1], (*pq)[1]
	*pq = (*pq)[0 : n-1]
	i := 1
	for {
		var a, b int = math.MaxInt32, math.MaxInt32
		var ind int

		if i*2 < n-1 {
			a = (*pq)[i*2]
		}
		if i*2+1 < n-1 {
			b = (*pq)[i*2+1]
		}

		if a < b {
			ind = i * 2
		} else {
			ind = i*2 + 1
		}

		if a == math.MaxInt32 && b == math.MaxInt32 {
			break
		}

		if (*pq)[ind] < (*pq)[i] {
			(*pq)[ind], (*pq)[i] = (*pq)[i], (*pq)[ind]
			i = ind
		} else {
			break
		}
	}
	return ret
}

type MaxPriorityQueue struct {
	q PriorityQueue
}

func (pq *MaxPriorityQueue) Push(x int) {
	pq.q.Push(-x)
}

func (pq *MaxPriorityQueue) Max() int {
	return -pq.q.Min()
}

func (pq *MaxPriorityQueue) Top() int {
	return -pq.q.Top()
}

func (pq MaxPriorityQueue) Len() int {
	return pq.q.Len()
}

type MedianFinder struct {
	n     int
	small MaxPriorityQueue
	large PriorityQueue
}

func (m *MedianFinder) Push(x int) {

	if m.n == 0 {
		m.small.Push(x)
		m.n++
		return
	}

	if m.small.Len() == m.large.Len() {
		if x < m.large.Top() {
			m.small.Push(x)
		} else {
			var min int = m.large.Min()
			m.small.Push(min)
			m.large.Push(x)
		}
	} else {
		if x < m.small.Top() {
			var max int = m.small.Max()
			m.large.Push(max)
			m.small.Push(x)
		} else {
			m.large.Push(x)
		}
	}
	m.n++
}

func (m *MedianFinder) Median() int {
	return m.small.Top()
}

func main() {
	n := 10000
	a := make([]int, n, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	var pq PriorityQueue = make([]int, 1, 5002)
	var maxPq PriorityQueue = make([]int, 1, 5002)

	solver := MedianFinder{0, MaxPriorityQueue{maxPq}, pq}

	sum := 0
	for i := 0; i < n; i++ {
		solver.Push(a[i])
		sum = (solver.Median() + sum) % 10000
	}

	fmt.Println(sum)
}
