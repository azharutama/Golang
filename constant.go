package main

import "fmt"

func main(){

	///constan adalah tipe data yang tidak bisa diubah

	const firstName string = "azhar"
	const lastName  = "utama"


	fmt.Println(firstName)
	fmt.Println(lastName)

	const(
		numberOne int = 1
		numberSecond = 2
	)

	fmt.Print(numberOne);
	fmt.Println(numberSecond)
}