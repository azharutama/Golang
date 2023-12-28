package main

import "fmt"

func main() {
	///pengecekan 1 variable

	name := "azhar"

	switch name {
	case "reza":
		fmt.Println("hello Azhar")
	case "azhar":
		fmt.Println("hello Azhar")
	default:
		fmt.Println("siapa?")

	}
///short statement
switch length:= len(name); length>4{
case true:
	fmt.Println("pas")
case false:
	fmt.Println("pas")
}


///switch tanpa kondisi
length:= len(name)
switch{
case length>10 :
	fmt.Println("terlalu panjang")
case length>5:
	fmt.Println("lumayan panjang")
default :
	fmt.Println("nice nice")

}
}