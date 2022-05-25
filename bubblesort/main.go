package main

import "fmt"

func main() {
	sort(SortBy{37, 5, 9, 11, 4, 33, 77, 3})
	sort(SortBy{37, 15, 90, 11, 4, 335, 77, 33})
}

func sort(arr SortBy) {
	for ind := 0; ind < arr.Len(); ind++ {
		for i := 0; i < arr.Len()-ind; i++ {
			if i == arr.Len()-1 {
				break
			}
			if arr.Less(i, i+1) {
				continue
			}
			arr.Swap(i, i+1)
		}
	}
	fmt.Println(arr)
}

type SortBy []int

func (a SortBy) Len() int           { return len(a) }
func (a SortBy) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortBy) Less(i, j int) bool { return a[i] < a[j] }
