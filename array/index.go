package main

import "fmt"

func main(){

	// integer
	var a[10]int
	fmt.Println(len(a)) // output is 10
	a[0] = 5
	fmt.Println(a) // output is [5 0 0 0 0 0 0 0 0 0]

	// boolean
	var b[5]bool 
	b[0] = true
	fmt.Println(b) // output is [true false false false false]

// string
	var c[10]string
	c[0] = "shazan"
	fmt.Println(c) // output is [shazan         ]


	// 2d array
	arr := [3][3]int{{5,5,5} ,{3,3,3}}
	fmt.Println(arr) // output is [[5 5 5] [3 3 3] [0 0 0]]


	// array 
	// ! fixed size - that is predictable 
	// ! memory optimization 
	// ! constant time action
}
