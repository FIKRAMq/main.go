package main

import "fmt"

func main() {
    bill := 275
    var tip float64

    if bill >= 50 && bill <= 300 {
        tip = float64(bill) * 0.1526
    } else {
        tip = float64(bill) * 0.2
    }

    total := float64(bill) + tip

    fmt.Printf("Tagihannya %d. tipnya %.2f, dan total nilainya %.2f", bill, tip, total)
}





