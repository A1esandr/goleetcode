package main

func twoSum(nums []int, target int) []int {
    for i, el := range nums {
        for ind, elem := range nums {
            if ind <= i {
                continue
            }
            if result := el + elem; result == target {
                return []int{i, ind}
            }
        }
    }
    return []int{}
}

func main() {
}
