package main

import "time"
import "fmt" 

type order struct {
	id int 
	amount float32
	status string
	createdAt time.Time

}

func (o *order) changeStatus(status string){
	o.status = status
	// *o.status = status 
	// we need not to use pointer in stuct .. because stuct by default do that
}

func (o *order) changeAmount(amount float32){
	o.amount = amount
}

func (o order) getId()int  {
	return o.id
}

func main() {
	fmt.Println("struct")
	myOrder := order{
		id: 1,
		amount: 50.20,
		status: "pending",
	}
	myOrder.createdAt = time.Now()

	fmt.Println("Order struct =" , myOrder) // output is -> Order struct = {1 50.2 Complete {13994398592634776448 286835 0x104f626e0}}

	// after change status with func and pointer..
	myOrder.changeStatus("confirmed")
	fmt.Println("Order struct after change the status =" , myOrder) // output -> Order struct after change the status = {1 50.2 confirmed {13994403584273761224 242584 0x100cea6e0}}


	// afteer change amount with func and pointer
	myOrder.changeAmount(4550)
	fmt.Println("Order struct after change the amount =" , myOrder) // output -> Order struct after change the amount = {1 4550 confirmed {13994403584273761224 242584 0x100cea6e0}}


	fmt.Println("myOrder id is =" , myOrder.getId()) // output -> 1



	// if you dont set any field, default value is zero value  
	
	//int default -> 0 , float default -> 0 , string default -> "" , bool default -> false
}
