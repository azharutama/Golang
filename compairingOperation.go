package main

import "fmt"
func main(){
	//unary operator

	var i = 2
	i++ //increment
	fmt.Println("hasil = ++", i)

	i--//decrement
	fmt.Println("hasil = --", i)

//operator perbandingan

 var name1 = "azhar"
 var name2 = "azhar"
 var resultTrue bool = name1==name2 
 var resultFalse bool = name1!=name2

 fmt.Println(1, resultTrue)
 fmt.Println(2, resultFalse)

 var number1 = 1
 var number2 = 2
 var lebihDari = number1 > number2
 var kurangDari = number1 < number2
 var lebihDariSamaDengan = number1 >= number2
 var kurangDariSamaDengan = number1 <= number2

 fmt.Println(">", lebihDari)
 fmt.Println("<", kurangDari)
 fmt.Println(">=", lebihDariSamaDengan)
 fmt.Println("<=", kurangDariSamaDengan)

}