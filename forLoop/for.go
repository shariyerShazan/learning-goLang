package main

import "fmt"

//  for is one and only for looping in goLang. no while or do while
func main(){
	// ! for loop 
	 for i := 0; i <= 10 ; i++ {
		 fmt.Println(i)
	 }

	// ! infinite loop
	// for {
	// 	fmt.Println("Infinite loop")
	// }


	// ! like while loop but writen by for keyword
     j := 1
	 for j < 10 {
		fmt.Println(j)
		j++
	 }


}
