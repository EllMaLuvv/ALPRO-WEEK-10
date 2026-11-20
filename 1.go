
package main
import "fmt"

func main() {
    var berat int
    fmt.Print("Berat parsel (gram): ")
    fmt.Scan(&berat)

    kg := berat / 1000
    sisa := berat % 1000
    biaya := kg * 10000
    tambahan := 0

    if berat > 10000 {
        tambahan = 0
    } else if sisa >= 500 {
        tambahan = sisa * 5
    } else if sisa > 0 {
        tambahan = sisa * 15
    }

    fmt.Printf("Detail berat: %d kg + %d gr\n", kg, sisa)
    fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biaya, tambahan)
    fmt.Printf("Total biaya: Rp. %d\n", biaya+tambahan)
}
