package main

import "fmt"
func main(){

	var uas = 80
  var uts = 70

	var lulusUts bool= uts >=80
	var lulusUas bool= uas >=80

	var lulusAND bool = lulusUts && lulusUas
	var lulusOR bool = lulusUts || lulusUas

	fmt.Println(lulusAND)
	fmt.Println(lulusOR)
}


