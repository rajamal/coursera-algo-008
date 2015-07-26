package main

import (
	"fmt"
)

func choosePivotLeft(arr []uint, l, r int) int {
	return l
}

func choosePivotRight(arr []uint, l, r int) int {
	return r - 1
}

func choosePivotMedian(arr []uint, l, r int) int {
	a := arr[l]
	b := arr[r-1]
	mid := (r-l-1)/2 + l
	c := arr[mid]
	if a < b {
		if b < c {
			return r - 1 // b
		} else {
			if c < a {
				return l // a
			} else {
				return mid // c
			}
		}
	} else {
		if a < c {
			return l // a
		} else {
			if b < c {
				return mid // c
			} else {
				return r - 1 // b
			}
		}
	}
}

func CountComparisons(arr []uint, l, r int, fun func([]uint, int, int) int) int64 {

	if r <= l+1 {
		return 0
	}

	pindex := fun(arr, l, r)
	arr[pindex], arr[l] = arr[l], arr[pindex]
	pivot := arr[l]

	i := l + 1
	for j := l + 1; j < r; j++ {
		if arr[j] < pivot {
			arr[j], arr[i] = arr[i], arr[j]
			i++
		}
	}

	arr[i-1], arr[l] = arr[l], arr[i-1]

	return CountComparisons(arr, l, i-1, fun) + CountComparisons(arr, i, r, fun) + int64(r-l-1)

}

func main() {
	n := 10000
	a := make([]uint, n, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	left := make([]uint, n, n)
	copy(left, a)
	fmt.Println(CountComparisons(left, 0, n, choosePivotLeft))

	right := make([]uint, n, n)
	copy(right, a)
	fmt.Println(CountComparisons(right, 0, n, choosePivotRight))

	median := make([]uint, n, n)
	copy(median, a)
	fmt.Println(CountComparisons(median, 0, n, choosePivotMedian))

}
