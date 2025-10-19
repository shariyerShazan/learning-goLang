package main

import "fmt"

// slice -> dynamic
// useful method

func main(){
	// uninitialized slice it nil
     var num[]int // empty array is null/nil
	 fmt.Println(num == nil) // output is true
	 fmt.Println(len(num)) // output 0


	 var number = make([]int , 0 , 5)

	 number = append(number, 1)
	 number = append(number, 2)
	 number = append(number, 3)
	 number = append(number, 4)
	 number = append(number, 5)
	 number = append(number, 6)
	 fmt.Println(cap(number ) ,"capacity")
	 fmt.Println(len(number) , "lenght")
}

