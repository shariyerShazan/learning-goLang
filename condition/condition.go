package main

import "fmt"

func main (){
	age := 20

	if age > 18 {
		fmt.Println("You are adult!")
	} else if age > 12{
		fmt.Println("You are teenager")
	} else if age > 0 {
		fmt.Println("You are child")
	}else {
		fmt.Println("Not a valid age!")
	}


	role := "admin" 
	isAuthenticated :=  true

	if role == "admin" && isAuthenticated {
		fmt.Println("Welcome admin")
	} else if role == "admin" {
        fmt.Println("well, you are admin. but not authenticated. please login first")
	} else {
		fmt.Println("Sorry: Access forbidden")
	}


	name := "shazan"
    password := 12930012
    var role1 = ""

    if name == "" || password == 0 || role1 == "" {
        fmt.Println("Something is missing")
    } else {
        fmt.Println("All good!")
    }
}
