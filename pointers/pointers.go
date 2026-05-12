package main

import "fmt"

// by value m
// func changeNum(num int) {
// 	num = 5
// 	fmt.Println("In change", num)
// }
// by reference

func changeNum(num *int) {
	*num = 5
	fmt.Println("In changeNum", *num)
}

func main() {
	num:= 1
	changeNum(&num)
	fmt.Println("In main", num)
}