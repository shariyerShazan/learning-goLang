package main

import "fmt"


func main(){
	m :=  make(map[string]string)

	m["Name"] = "Shazan" 
	m["Title"] = "Anna"
	fmt.Println(m["Name"]  , m["Title"] , m["NotExist"]) //output is ->  Shazan Anna ______
}
