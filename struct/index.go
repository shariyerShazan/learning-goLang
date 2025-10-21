package main

import "time"
import "fmt" 

type order struct {
	id int 
	amount float32
	status string
	createdAt time.Time

}


func main() {
	fmt.Println("struct")
	myOrder := order{
		id: 1,
		amount: 50.20,
		status: "Complete",
	}
	myOrder.createdAt = time.Now()

	fmt.Println("Order struct =" , myOrder) // output is -> Order struct = {1 50.2 Complete {13994398592634776448 286835 0x104f626e0}}
}
