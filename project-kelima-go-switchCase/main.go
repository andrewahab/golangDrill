package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var nama string

	fmt.Println("Masukan Nama:")
	fmt.Scanln(&nama)

	score := rand.Intn(101)

	switch {
	case score > 80:
		fmt.Println("Selamat", nama, ", anda sangat beruntung, nilai anda -->", score)
	case score <= 80 && score > 60:
		fmt.Println("Selamat", nama, ", anda beruntung, nilai anda -->", score)
	case score <= 60 && score > 40:
		fmt.Println("Sayang sekali", nama, ", anda kurang beruntung, nilai anda -->", score)
	case score <= 40:
		fmt.Println("Wah", nama, ", anda sangat tidak beruntung, nilai anda -->", score)
	}

}