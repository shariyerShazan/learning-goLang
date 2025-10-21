package main

import "fmt"

func add(a int , b int) int {
	return a + b
}

func multiReturn()(string , string , string , bool){
	return  "javaScript" , "typeScript" , "goLang" , true
}
func main(){
	 result := add(3,5)
	 fmt.Println(result)
	 
	 lang1 , lang2 , lang3 , _ := multiReturn()
	 fmt.Println(lang1 , lang2 , lang3)
}