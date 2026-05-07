package main

// iterate over a range of numbers

import "fmt"


func main(){
	//nums:= []int{1, 2, 3, 4, 5}

	// for i:=0; i<len(nums); i++{
	// 	fmt.Println(nums[i])
	// }

	// for _, num := range nums{           // _ is index
	// 	fmt.Println(num)
	// }




	// range on maps

	// m:= map[string]int{
	// 	"one": 1,
	// 	"two": 2,
	// 	"three": 3,
	// }

	// for k, v:= range m{
	// fmt.Println(k,v)
	//}


	// range on strings

	s:= "hello world"

	for i, c:= range s{
		//fmt.Println(i,c);  // c is unicode point rune  , // i - starting byte of rune
		fmt.Println(i, string(c));  // convert rune to string
	}

	
}