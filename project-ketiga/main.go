package main

import "fmt"

func main() {
	var people []string = []string{"Walt", "Jesse", "Skyler", "Saul"}

	fmt.Println("Panjang slice people", len(people))

	var newpeople = append(people, "Hank", "Marie")

	fmt.Println("Jumlah baru panjang slice people", len(newpeople))

}