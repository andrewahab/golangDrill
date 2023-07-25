package main

import (
	"fmt"
	"strconv"
)

func main() {

	var nama string
	fmt.Println("Masukan Nama:")
	fmt.Scanln(&nama)

	var umurStr string
	fmt.Println("Masukan Umur:")
	fmt.Scanln(&umurStr)

	umur, err := strconv.Atoi(umurStr)
	if err != nil {
		fmt.Println("ERROR: Umur bukan angka")
		return
	}

	switch {
		case umur < 0 || umur > 100:
		fmt.Println("ERROR")
		case umur > 18:
		fmt.Println("Congratulation", nama, ", you're", umur, ", you may enter!")
		case umur <= 18:
		fmt.Println("I'm sorry", nama, ", you're", umur, ", children is not allowed to enter!")
		
	
	}
}