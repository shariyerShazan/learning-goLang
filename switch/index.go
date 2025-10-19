package main

import (
	"fmt"
	"time"
)

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

	 switch time.Now().Weekday(){
	    case time.Sunday , time.Friday :
		    fmt.Println("It's weekend")
		default : 
		    fmt.Println("Work day")
	 }

}