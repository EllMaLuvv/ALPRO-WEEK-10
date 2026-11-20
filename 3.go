package main

import "fmt"

func main() {
	var b int
	fmt.Print("Bilangan: ")
	fmt.Scan(&b)

	fmt.Print("Faktor:")
	jumlahFaktor := 0
	for faktor := 1; faktor <= b; faktor++ {
		if b%faktor == 0 {
			fmt.Printf(" %d", faktor)
			jumlahFaktor++
		}
	}
	println()
	if b > 0 {
		if jumlahFaktor == 2 {
			fmt.Println("Prima: true")
		} else {
			fmt.Println("Prima: false")
		}
	}
}
