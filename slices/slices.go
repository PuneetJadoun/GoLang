package main

import (
	"fmt"
	"slices"
)

// slice - dynamic
// most used construct in Go, built on top of arrays
// useful methods

func main(){


	//uninitialised slice is nil and has length and capacity of 0

	// var nums []int

	// fmt.Println(nums) // []
	// fmt.Println(nums == nil) // true
	// fmt.Println(len(nums))

	 
	// var nums = make([]int, 2, 5) // length 2, capacity 5
	// fmt.Println(len(nums)) // 2
	// fmt.Println(cap(nums)) // 5
	// fmt.Println(nums) // [0 0]

	// nums:= []int{}
	// nums[0] = 1 // panic: runtime error: index out of range [0] with length 0
	// fmt.Println(nums) // []

	// copy function 

	// var nums = make([]int, 0, 5)
	// var nums1 = make([]int, len(nums))

	// nums = append(nums, 1)


	// copy(nums, nums1)



	// sclice operator

	// var nums = []int{1, 2, 3, 4, 5}

	// fmt.Println(nums[1:4])


	// slice 
	var nums = []int{1, 2, 3, 4, 5}
	var nums1 = nums[1:4] // creates a new slice that references the same underlying array
	
	fmt.Println(slices.Equal(nums, nums1)) // true



	// 2d slices

	var matrix = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	fmt.Println(matrix)

}