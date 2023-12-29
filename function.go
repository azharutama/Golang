package main

import "fmt"

func sayHello() {
	fmt.Println("hello")
}

func countNumber(a int,b int){
	fmt.Println(a*b)
}

func showName(namaDepan string, namaBelakang string){
fmt.Println(namaDepan,namaBelakang)
}

func countLenName(name string){
	fmt.Println(len(name))
}
//fungsi return value
func getHello(name string) string{

	hello :=  "hello " + name
	return hello
	
}

func main() {
	//panggil fungsi biasa
	sayHello()
	countNumber(3,3)
	showName("azhar","utama")
	countLenName("azhar utama")

	//panggil fungsi return
	result:=getHello("azhar")
	fmt.Println(result)
	///cara lain fungsi return
	fmt.Println(getHello("utama"))


}

