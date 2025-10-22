package main

import "fmt"

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int) // เก็บค่ากับ index
	for i, num := range nums {
		diff := target - num
		if j, ok := numMap[diff]; ok {
			return []int{j, i}
		}
		numMap[num] = i
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9

	result := twoSum(nums, target)
	if result != nil {
		fmt.Printf("Index: %v (values: %d + %d = %d)\n",
			result, nums[result[0]], nums[result[1]], target)
	} else {
		fmt.Println("No match found.")
	}
}
