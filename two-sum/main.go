package main

import "fmt"

func twoSum(nums []int, target int) []int {
    numMap := make(map[int]int)
    for i, num := range nums {
        complement := target - num
        if j, ok := numMap[complement]; ok {
            return []int{j, i}
        }
        numMap[num] = i
    }
    return nil
}

func main() {
	arr := []int{3,3}

	fmt.Printf("%v", twoSum(arr, 6))
}
