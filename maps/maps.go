package main

import (
	"fmt"
	"maps"
)

// maps -- hash , objects , dict

func main(){
	// creating map

	// m:= make(map[string]string)

	// setting ele

	// m["name"] = "golang"

	// // get ele

	// fmt.Println(m["name"])

	// // if key value not present in map then it will return zero value of that type

	// // len of map

	// fmt.Println(len(m))

	// // delete ele	

	// delete(m,"name")

	// fmt.Println(m["name"])

	// clear(m)


	// m:= map[string]int{
	// 	"age": 12,
	// 	"year": 2024,
	// }

	// fmt.Println(m["age"])
	// fmt.Println(m["year"])


	// count, ok := m["price"]
    // fmt.Println(count)

	// if ok{
	// 	fmt.Println("price is present in map")
	// }else{
	// 	fmt.Println("price is not present in map")
	// }


	m1 := map[string]int{
		"age": 12,
		"year": 2024,
	}

	m2 := map[string]int{
		"age": 12,
		"year": 2024,
	}

	// check if m1 and m2 are equal

	fmt.Println(maps.Equal(m1, m2))
	

}