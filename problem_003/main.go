package main

import "fmt"

// This problem was asked by Stripe.
// Given an array of integers, find the first missing positive integer in linear time and constant space.
// In other words, find the lowest positive integer that does not exist in the array. The array can contain duplicates and negative numbers as well.
// For example, the input [3, 4, -1, 1] should give 2. The input [1, 2, 0] should give 3.
// You can modify the input array in-place.

func main() {
	input := []int{3, 4, -1, 1}
	//input := []int{3, 4, -1, -5, 1, 10, 2, 3, 4, 3, 2, 1, 7, 9, -2, -1, 0, 5, 6, 10}
	fmt.Println(lowestMissingPositive(input))
}

func lowestMissingPositive(nums []int) int {
	n := len(nums)
	for i := 0; i < n; i++ {
		for nums[i] > 0 && nums[i] <= n && nums[nums[i]-1] != nums[i] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}

	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}

	return n + 1
}
