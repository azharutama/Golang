package main

import "fmt"

func main() {
	number := 1
	for number <= 10 {
		fmt.Println("perulangan ke", number)
		number++
	}
	fmt.Println("selesai")

	//init statement = awal perulangan
	for id:=1; id <= 10; id++{
		fmt.Println(id)
		
	}
	fmt.Println("selesai")

	///for range

	//manual
	names:= []string {
		"azhar","reza","yadi",
	}
	for i:= 0 ; i < len(names);i++{
		fmt.Println(names[i])
	}

	//pakai range
	for index, name := range names{
		fmt.Println("index", index ,"=" ,name)
	}

	//tanpa index
	for _, name := range names{
		fmt.Println(name)
	}

	numbers := []int{
		1,2,3,4,5, 
	}

	for _, no := range numbers{
		fmt.Println(no)
	}
//menggunakan map
	id := map[string]int{
		"no":1,
		"name": 2,
	}
 	for index,idBaru := range id{
		fmt.Println(index,idBaru)
	} 

	slice1:= numbers[2:]
	for _, slice:= range slice1{
		fmt.Println("nomor",slice)
	}

}