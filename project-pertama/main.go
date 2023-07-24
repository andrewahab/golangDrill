package main

import "fmt"

func main() {
	// var myNum int32 = 50
	// var myNum2 float32 = 51
	// var myNumStr string = "50"

	// fmt.Println(myNum)
	// fmt.Println(myNum2)
	// fmt.Println(myNumStr)

	// var x int32 = 5
	// var y int32 = 10
	// var z int32 = (x + y)

	// fmt.Println(z)

	var nama, alamat string

	fmt.Println("Masukan nama:")
	fmt.Scanln(&nama)

	fmt.Println("Masukan alamat:")
	fmt.Scanln(&alamat)

	fmt.Println("Hello", nama)
	fmt.Println("Alamat:", alamat)
	fmt.Println(nama, "tinggal di", alamat)

	var people []string = []string{"Walt", "Jesse", "Skyler", "Saul"}

	fmt.Println("Panjang slice people", len(people))

	var newpeople = append(people, "Hank", "Marie")

	fmt.Println("Jumlah baru panjang slice people", len(newpeople))

	// identitas := map[string]string{
	// 	"name": "Hank",
	// 	"gender": "M",

	// }
	// fmt.Println(identitas)

}
