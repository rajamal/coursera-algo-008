package main

import (
	"fmt"
)

func CountInversions(arr []uint) int64 {

	n := len(arr)
	if n == 1 {
		return 0
	}

	mid := len(arr) / 2
	left := make([]uint, mid)
	right := make([]uint, n-mid)

	copy(left, arr[0:mid])
	copy(right, arr[mid:])

	a := CountInversions(left)
	b := CountInversions(right)

	c := int64(0)
	i := 0
	j := 0
	for i < mid && j < (n-mid) {
		if left[i] <= right[j] {
			arr[i+j] = left[i]
			i++
		} else {
			arr[i+j] = right[j]
			j++
			c += int64(mid - i)
		}
	}

	for i < mid {
		arr[i+j] = left[i]
		i++
	}

	for j < n-mid {
		arr[i+j] = right[j]
		j++
	}
	return a + b + c

}

func main() {
	n := 100000
	a := make([]uint, n, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	fmt.Println(CountInversions(a))
}
