package main

import "fmt"

func main(){
	
number:= [...]int{1,2,3,4,5,6,7,8,} 

slice1 := number [4:5]
fmt.Println(slice1)

slice2 := number [:5]
fmt.Println(slice2)

slice3 := number [:]
fmt.Println(len(slice3))

days := []int {
	1,2,3,4,5,6,7,
}

daySlice1:= days [1:]
daySlice1[0]= 2
daySlice1[3]= 2
fmt.Println(days)


//append
daySlice2 := append(daySlice1, 3)
fmt.Println(daySlice1)
fmt.Println(daySlice2)
fmt.Println(days)

//make
newSlice := make([]int, 2,5)
newSlice[0]=1
newSlice[1]=2

fmt.Println(newSlice)
fmt.Println(len(newSlice))
fmt.Println(cap(newSlice))

newSlice2 := append(newSlice, 3)
fmt.Println(newSlice2)


//copy
fromSlice := days[:]
toSlice := make([]int , len(fromSlice) , cap(fromSlice))
copy(toSlice, fromSlice)
fmt.Println(fromSlice)
fmt.Println(toSlice)
}