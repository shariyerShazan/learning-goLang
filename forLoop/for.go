package main

import "fmt"

//  for is one and only for looping in goLang. no while or do while
func main(){
	// ! for loop 
	fmt.Println("Foor loop")
	 for i := 0; i <= 10 ; i++ {
		 fmt.Println(i)
	 }

	// ! infinite loop
	// for {
	// 	fmt.Println("Infinite loop")
	// }


	// ! like while loop but writen by for keyword
	fmt.Println("While but written like for")
     j := 0
	 for j <= 10 {
		// ?  break  // for exit from the loop
		// ? continue // for exit for the current value of index

		fmt.Println(j)
		j++
	 }
 
	// ! range looop
	fmt.Println("Rnage loop")
	for x:= range 10 {
		fmt.Println(x)
	}

}
