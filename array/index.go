package main

import "fmt"

func main(){
	var a[10]int
	fmt.Println(len(a)) // output is 10
	a[0] = 5
	fmt.Println(a) // output is [5 0 0 0 0 0 0 0 0 0]

	var b[5]bool 
	b[0] = true
	fmt.Println(b) // output is [true false false false false]


	var c[10]string
	c[0] = "shazan"
	fmt.Println(c) // output is [shazan         ]
}
