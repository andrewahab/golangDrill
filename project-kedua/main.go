package main

import "fmt"

func main() {
	var nama, alamat string

	fmt.Println("Masukan nama:")
	fmt.Scanln(&nama)

	fmt.Println("Masukan alamat:")
	fmt.Scanln(&alamat)

	fmt.Println("Hello", nama)
	fmt.Println("Alamat:", alamat)
	fmt.Println(nama, "tinggal di", alamat)

}