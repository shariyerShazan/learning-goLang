package main

import "fmt"

func main(){

	// integer array
	 nums := []int{1,5,10}
	 sum := 0
	 for index , num := range nums {
		sum += num
		fmt.Println(index , "No of value is =" , num)
	 }
	 fmt.Println( "Total sum = ", sum)



	//  string array
	lifeStatus := map[string]string{"name" : "shazan" , "wife" : "priya"}
	for key , value := range lifeStatus {
		fmt.Println(key , value)
	}
}
