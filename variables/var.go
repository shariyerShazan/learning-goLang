package main

import "fmt"


func main(){

	// type defined
	var name string = "shazan";
	fmt.Println(name);

	// no need to define type & auto infer type
   var keyword = "shariyer";
   fmt.Println(keyword);


   //  int value
   var age int = 22;
   fmt.Println(age);


	//short way to declear variable.
	fullName := "Shariyer Shazan" ;
	fmt.Println(fullName);

// string with var
	var fatherName string
	  fatherName = "Khairul"
	  fmt.Println(fatherName)



	//   float with var
	  var price float32 = 25.08 
	  fmt.Println(price)


	//   float without var
	price2 := 200.20
	fmt.Println(price2)
}