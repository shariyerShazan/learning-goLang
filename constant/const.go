package main

import "fmt"

const fullName = "shariyer shazan" // we can declear const outside of main function
var age2 = 22 // we can declear var outside of main function

// text := "sms" //! we can't declear shortform or declearation outside of the main function

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

	fmt.Println(fullName)
	fmt.Println(age2)
}
