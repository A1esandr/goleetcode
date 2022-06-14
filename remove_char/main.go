package main

import "fmt"

func main() {
	fmt.Println(remove([]int{1, 1, 2}, 1))
	fmt.Println(remove([]int{3, 1, 1, 2}, 1))
	fmt.Println(remove([]int{3, 1, 2}, 1))
	fmt.Println(remove([]int{3, 2, 1, 1}, 1))
	fmt.Println(remove([]int{1, 1, 4, 1, 1, 1, 3, 2, 1, 1}, 1))
	fmt.Println(remove([]int{1, 1, 4, 1, 1, 1, 3, 1, 1, 2, 1, 1}, 1))
}

func remove(numbers []int, number int) []int {
	inc := 1
	last := 0
	for i := 0; i < len(numbers); {
		if numbers[i] != number {
			i++
			if i+last == len(numbers)-1 {
				if numbers[i] == number {
					last++
				}
				break
			}
			if inc > 1 {
				inc--
			}
			continue
		}
		if i == 0 {
			numbers = numbers[1:]
			continue
		}
		if i == len(numbers)-1 {
			numbers = numbers[:i]
			break
		}
		if i+inc == len(numbers) {
			break
		}
		numbers[i] = numbers[i+inc]
		numbers[i+inc] = number
		last++
		if i+last == len(numbers)-1 {
			if numbers[i] == number {
				last++
			}
			break
		}
		inc++
	}
	if last > 0 {
		numbers = numbers[:len(numbers)-last]
	}
	return numbers
}
