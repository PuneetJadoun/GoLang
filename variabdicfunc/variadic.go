package main


import "fmt"

func sum(nums ...int) int{  // if want any type use interface{} instead of int
	total := 0

	for _, num := range nums{
		total = total + num
	}
	return total
}


func main(){
	//fmt.Println(1, 2, 3, 4 , 5 , "h")

	nums:= []int{1, 2, 3, 4, 5}
	//result := sum(1, 2, 3, 4)
	result := sum(nums...)  // if want any type use interface{} instead of int	
	fmt.Println(result)
} 