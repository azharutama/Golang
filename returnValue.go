package main

import "fmt"

func getFullName() (string, string) {
	return "azhar", "utama"
}

func getFullName2() (string, string) {
	return "azhar", "utama"
}


func getCompleteName()(firstName3, middleName3, lastName3 string){
firstName3 = "Muhammad"
middleName3= "azhar"
lastName3 = "utama"

return firstName3,middleName3,lastName3
}

func getDataMahasiswa()(Nama string, NIM int, ipk float32){
	Nama= "azhar"
	NIM = 223040077
	ipk = 3.48
	
	return Nama, NIM,ipk
	}

func main() {
	firstName, lastName := getFullName()
	fmt.Println(firstName, lastName)

	firstName2, _ := getFullName()
	fmt.Println(firstName2)

a,b,c:= getCompleteName()
	fmt.Println(a,b,c)

	nama,nim,ipk := getDataMahasiswa()
	fmt.Println(nama,nim,ipk)
	//print per field
	fmt.Println("nama =",nama)
	fmt.Println("Nim =", nim)
	fmt.Println("IPK =",ipk)


}