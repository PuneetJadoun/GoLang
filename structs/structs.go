package main

import (
	"fmt"
	"time"
)

// order struct


type customer struct{
	name string
	phone string
}
type order struct{
	id string
	amount float32
	status string
	createdAt time.Time  // nanosecond precision
	customer   // struct embeding
}

// receiver type

// func (o *order) changeStatus(status string){
// 	o.status = status
// }

// func newOrder(id string, amount float32, status string) *order {
// 	myOrder:= order{
// 		id: id,
// 		amount: amount,
// 		status: status,
// 	}

// 	return &myOrder
// }




func main(){
	// if you dont set any fields, they will be set to their zero value
	// myOrder := order{
	// 	id: "1",
	// 	amount: 50.00,
	// 	status: "pending",
	// }

	// myOrder.createdAt = time.Now()
	// myOrder.changeStatus("confirmed")

	// fmt.Println("order struct: ", myOrder)	
	
	
	//myOrder := newOrder("1", 30, "received")
	//fmt.Println("order struct: ", myOrder)


	//language:= struct{
	// 	name string
	// 	isGood bool
	// }{"Go", true}

	//fmt.Println("language struct: ", language)


	newCustomer := customer{
		name: "John Doe",
		phone: "1234567890",
	}

	newOrder := order{
		id: "1",
		amount: 50.00,
		status: "pending",
		customer: newCustomer,
	}

	fmt.Println("order struct: ", newOrder)
}	