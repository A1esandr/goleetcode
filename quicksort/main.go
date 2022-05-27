package main

import "fmt"

func main() {
	arr := []int{77, 32, 44, 9, 101, 1, 32, 55, 6}
	quicksort(arr)
	fmt.Println(arr)
}

func quicksort(arr []int) {
	if len(arr) < 2 {
		fmt.Println(arr)
		fmt.Println("return")
		return
	}
	fmt.Println(arr)
	low, high := 0, len(arr)-1
	for i, _ := range arr {
		if arr[i] < arr[high] {
			arr[i], arr[low] = arr[low], arr[i]
			low++
		}
	}

	arr[low], arr[high] = arr[high], arr[low]

	fmt.Println(arr)
	fmt.Println(low)

	quicksort(arr[:low])
	quicksort(arr[low+1:])
}
