package main

import "fmt"

func sum(nums ...int)int {
	total := 0
	for _ , num:= range nums {
         total += num
	}
	return  total
}

func main(){
	fmt.Println(1,2,3,4,5,6 , "shazan")

	result := sum(1,2,3,4,5,6)
	fmt.Println(result)

	nums2 := []int{1, 2, 3, 4, 5, 6}
	result2 := sum(nums2...)
	fmt.Println("result two by variadic= ", result2)
}