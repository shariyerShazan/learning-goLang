package main

import "fmt"

func changeNum(num int){
	num = 5
	fmt.Println("Changed num =" , num)
}

func changeNum2(num *int){
	*num = 5
	fmt.Println("Changed num =" , *num)
}
func main(){
  num := 10 
    fmt.Println("Before change num in main =" , num) 
  changeNum(num)
  fmt.Println("After change num in main =" , num) // not change.. because we just send copy of the num on changeNum funciton. we need to send the loaction of the variable.. then we can changhe the value... so sent varibale location not it's copy..

  num2 := 100 
  fmt.Println("Memory adress of num2 =", &num2)
    fmt.Println("Before change num2 in main =" , num2)
  changeNum2(&num2)
  fmt.Println("After change num2 in main =" , num2)

}
