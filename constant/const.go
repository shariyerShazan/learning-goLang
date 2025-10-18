package main

import "fmt"

func main(){
	var name = "shazan" 
	fmt.Println(name)

	// you can reInitial var
	name = "priya"
	fmt.Println(name)

	// you cant reInitial const
	const age = 20
	// age = 200 //! its will error
	fmt.Println(age)
}
