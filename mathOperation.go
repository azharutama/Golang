package main

import "fmt"

func main(){

	///matematika dasar
	var a = 10
	var b = 20
	var c = a*b
 	var d = a+b
	var e = a-b
	var f = a/b
	var g = a%b
	
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)

	///augmented assignment
	var h = 10
			h += 20
			fmt.Println("hasil +=", h)
			h -= 20
			fmt.Println("hasil -=", h)
			h *= 20
			fmt.Println("hasil *=", h)
			h /= 20
			fmt.Println("hasil /=", h)
			h %= 20
			fmt.Println("hasil %=", h)


			
	
			

}