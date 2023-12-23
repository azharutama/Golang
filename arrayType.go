package main

import "fmt"
func main(){
	var names [3]string
	names[0] = "muhammad"
	names[1] = "azhar"
	names[2] = "utama"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	///secara langsung
	var nama =[3]string{
		"muhammad","azhar","utama",
	}

	fmt.Println(nama)
	
	//tanpa memberi nilai

	var mahasiswa =[...]string{
		"muhammad","azhar","utama",
	}

	fmt.Println(mahasiswa)
	fmt.Println(len(mahasiswa))//menghitung nilai array
	mahasiswa[0]="reza"	//mengganti nilai array
	fmt.Println(mahasiswa)

	var number =[...]int{
		1,2,3,
	}

	fmt.Println(number)
	fmt.Println(len(number))

	number[2]=30
	fmt.Println(number)

}