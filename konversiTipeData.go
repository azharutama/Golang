package main

import "fmt"

func main(){
	//nilai integer
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai64)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)
	

	///nilai 
	 var name string = "azhar"
var a = name[0]////dalam bentuk integer
var aString = string(a) ///convert dalam bentuk string


fmt.Println(name)
fmt.Println(a)
fmt.Println(aString)
} 
