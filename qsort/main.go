package main

import (
	"fmt"
)

func main() {
	arr := []int{3, 2, 1, 3, 5, 5, 4}
	quicksort(arr)
	fmt.Println(arr)
}

func quicksort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	// pivot := rand.Int() % len(a)

	// a[pivot], a[right] = a[right], a[pivot]

	for i, _ := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	fmt.Println(a)
	fmt.Println(left)

	quicksort(a[:left])
	quicksort(a[left+1:])

	return a
}
