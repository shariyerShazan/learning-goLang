package main

import (
	"fmt"
	"maps"
)


func main(){
	m :=  make(map[string]string) // (map[key type] value type)

	m["Name"] = "Shazan" 
	m["Title"] = "Anna"
	fmt.Println(m["Name"]  , m["Title"] , m["NotExist"]) //output is ->  Shazan Anna ______


	// already have data then initialized like ->
     m2 := map[string]int{"Price" : 40 , "Quantity" : 12} // output -> map[Price:40 Quantity:12]
	 m3 := map[string]int{"Price" : 40 , "Quantity" : 12} 
	 fmt.Println(m2)

	 value , exist := m2["Price"]
	 fmt.Println(value) // output -> 40
	 if exist {
		fmt.Println(exist ,"Is exist") // output -> true Is exist
	 } else {
		fmt.Println(exist ,"Is not exist")
	 }

	 fmt.Println(maps.Equal(m2 , m3)) // output -> true
}
