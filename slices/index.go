package main

import (
	"fmt"
	"slices"
)

// slice -> dynamic
// useful method

func main(){
	// uninitialized slice it nil
     var num[]int // empty array is null/nil
	 fmt.Println(num == nil) // output is true
	 fmt.Println(len(num)) // output 0


	 var number = make([]int , 0 , 5) // it's takes (_ , default length , defalt capacity)

	 number = append(number, 1)
	 number = append(number, 2)
	 number = append(number, 3)
	 number = append(number, 4)
	 number = append(number, 5)
	 number = append(number, 6)

	 fmt.Println(len(number) , "lenght") //length of slice
	 fmt.Println(cap(number ) ,"capacity") // capacity of slice



    //  slice start to end
	 var num2 = []int{1,2,3,4,5}
	 fmt.Println(num2[0:3]) // output is [1 2 3]
	 fmt.Println(num2[:3]) // output is [1 2 3]


	//  compare slice
	var num3 = []int{1, 2, 3}
	var num4 = []int{1, 2, 3}
	var num5 = []int{5, 5, 5}

	fmt.Println(slices.Equal(num3 , num4)) // output -> true
	 copy(num3 , num5)
	 fmt.Println(num3) // output -> [5 5 5]

}

