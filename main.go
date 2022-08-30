package main

import (
	"fmt"
	"multiple-return/name"
	"multiple-return/utils"
)

func main() {
	var namaDepan, namaBelakang, _ = name.GetName("Regi", "Rianto", 21)

	fmt.Println(namaDepan)
	fmt.Println(namaBelakang)

	scores := []int{90, 100, 20}
	hasil := utils.Sum(scores)

	fmt.Println(hasil)

}
