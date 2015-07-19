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
	ind := 0
	for i < mid && j < (n-mid) {
		if left[i] <= right[j] {
			arr[ind] = left[i]
			ind++
			i++
		} else {
			arr[ind] = right[j]
			j++
			ind++
			c += int64(mid - i)
		}
	}

	for i < mid {
		arr[ind] = left[i]
		ind++
		i++
	}

	for j < n-mid {
		arr[ind] = right[j]
		ind++
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
