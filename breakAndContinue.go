package main

import "fmt"

func main() {
	//break
	//menghentikan perulangan
	for i := 1; i <= 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println("no ", i)
	}

	//continue
	//untuk melanjutkan perulanan, jadi jika kondisi terpenuhi maka akan lanjut ke loop selanjutnya dan nilai dengan kondisi yang terpenuhi tidak di tampilkan 
	for i := 1; i <= 10; i++ {
		if i % 2 == 0 {
			continue
		}
		fmt.Println("no ", i)
	}
}