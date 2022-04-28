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

func twoSumV2(nums []int, target int) []int {
    for i, el := range nums {
        for ind:=i+1;ind<len(nums);ind++ {
            if result := el + nums[ind]; result == target {
                return []int{i, ind}
            }
        }
    }
    return []int{}
}

func main() {
}
