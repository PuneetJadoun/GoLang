package main

import "fmt"	

import "time"


func main(){
	// switch statement

	day := "Monday"			
	switch day {
	case "Monday":
		fmt.Println("It's Monday.")
	case "Tuesday":
		fmt.Println("It's Tuesday.")
	default:
		fmt.Println("It's another day.")
	}


	// multiple switch case

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("It's a weekday.")
	}

	// type switch

	whoAmI := func(i interface{}) {
		switch v := i.(type) {
		case int:
			fmt.Printf("I am an integer: %d\n", v)
		case string:
			fmt.Printf("I am a string: %s\n", v)
		default:
			fmt.Printf("I am of a different type: %T\n", v)

		}
	}	

	whoAmI(42)
}