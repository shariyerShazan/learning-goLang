package main

import "fmt"

func counter() func() int{
	count := 0

	return func() int {
		count ++
		return count
	}
}
func main(){
	increment := counter()

	for i:= 1 ; i< 100 ; i++ {
		increment()
	}
	fmt.Println(increment())
}
