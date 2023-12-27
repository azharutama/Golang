package main

import "fmt"

func main() {

	var mahasiswa map[string]string = map[string]string{}
	mahasiswa["name"] = "azhar"
	mahasiswa["NIM"] = "223040077"

	///cara cepat

	student := map[string]string{
		"name": "azhar",
		"NIM":  "223040077",

	}

	
	fmt.Println(student["name"])
	fmt.Println(student["NIM"])
	fmt.Println(student)


	//map program
	book:=make(map[string]string)//makeMap
	book["name"] = "tes"//insertData
	book["forDeleteTest"] = "Delete"

	delete(book,"forDeleteTest")//deleteData

	fmt.Println(book)//getData



}