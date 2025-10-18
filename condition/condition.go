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
}
