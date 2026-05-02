package main

import "fmt"

// numbered sequence of elements of the same type
func main(){
	// arrays 

	// zeros value  int - 0, string - "", bool - false
	// var nums [4]int

	// nums[0] = 1
	
	// fmt.Println(nums[0])

	// fmt.Println(nums)

	// fmt.Println(len(nums))

	// nums:= [4]int{1, 2, 3, 4}

	//fmt.Println(nums)


	// 2d arrays 
	nums := [2][2]int{{2,3}, {4,5}}
	fmt.Println(nums)


	// fixed size arrays are not flexible, we can use slices instead
	
}