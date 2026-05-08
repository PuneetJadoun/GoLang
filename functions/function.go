package main

import "fmt"

// func add(a int, b int) int {
// 	return a + b
// }


// func getLanguages() (string, string, string) {
// 	return "Go", "Python", "JavaScript"
// }

// func getLanguages() (string, string, bool) {
// 	return "Go", "Python", true
// }


// func processIt(fn func(a int) int){
// 	fn(1)
// }


func processIt() func(a int)int{
	return func (a int) int {
		return 4
	}
}


func main(){
	//result := add(5, 3)
	//println(result)

	//fmt.Println(getLanguages())

	//lang1, lang2, _ := getLanguages()
	//fmt.Println(lang1)
	//fmt.Println(lang2)

	// fn:= func(a int) int {
	// 	return 2
	// }

	//processIt(fn)

	fn:= processIt()
	fmt.Println(fn(1))
	
}