package main

import (
	"fmt"
	"time"
)


// number switch
func main(){
	i := 2
	switch i {
	  case 1 :
		fmt.Println("Case one")
		  // ! no need to write break // break
	  case 2 : 
	     fmt.Println("Case two")
	  case 3 : 
	     fmt.Println("Case three")
	  case 4 : 
	     fmt.Println("Case four")
	  default: 
	     fmt.Println("Wrong input")
	 }


	//  date with switch
	 switch time.Now().Weekday(){
	    case time.Sunday , time.Friday :
		    fmt.Println("It's weekend")
		default : 
		    fmt.Println("Work day")
	 }


	//  data type check by switch and function
	dataType := func(i interface{}){
		switch TYPE := i.(type){
		case int : fmt.Println("It's a integer")
		case string : fmt.Println("It's a string")
		case bool : fmt.Println("It's a boolan")
		default : fmt.Println("Invalid type! please try again" , TYPE)
		}
	}
	dataType("23")
}
