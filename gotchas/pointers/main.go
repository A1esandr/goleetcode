package main

import "fmt"

func main() {
	printGot()
}

func printGot() {
	result := got()
	for _, r := range result {
		fmt.Println(*r)
	}
	fmt.Println(got())
}

// problem
// variable v is always the same in loop
// it has exactly the one memory address
func got() []*int {
	a := []int{1, 2, 3, 4}
	result := make([]*int, 0)
	for _, v := range a {
		result = append(result, &v)
	}
	return result
}
