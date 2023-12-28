package main

import "fmt"

func main() {
	name := "reza"

	if name == "azhar" {
		fmt.Println("hello", name)
	}else if name == "reza" {
		fmt.Println("reza")
	}else{
		fmt.Println("siapa???")
	}

	///short statement
 ///name diambil dari name = reza
 ///dilakukan ketika ingin melakukan pengecekan sederhana
	if length:=len(name) ;length>5 {
		fmt.Println("terlalu panjang")
	}else{
		fmt.Println("pas")
	}
}