package main

import (
	"fmt"
	"sort"
)

type Int64Arr []int64

func (a Int64Arr) Len() int {
	return len(a)
}

func (a Int64Arr) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a Int64Arr) Less(i, j int) bool {
	return a[i] < a[j]
}

func main() {
	n := 1000000
	a := make(Int64Arr, n, n)
	hash := make(map[int64]bool)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
		hash[a[i]] = true
	}
	sort.Sort(a)

	var t int64
	count := 0
	for t = -10000; t <= 10000; t++ {
		fmt.Println("Solving for ", t)
		for i := 0; i < n; i++ {
			if a[i] < t-a[i] {
				if hash[t-a[i]] == true {
					count++
					break
				}
			}
			if a[i] > 10000 {
				break
			}
		}
	}
	fmt.Println(count)

}
