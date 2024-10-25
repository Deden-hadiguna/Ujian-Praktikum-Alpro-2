//2311102286
//Deden Hadiguna

package main

import "fmt"

func isPerfect(num int) bool {
	sum := 0
	for i := 1; i <= num/2; i++ {
		if num%i == 0 {
			sum += i
		}
	}
	return sum == num
}

func main() {
	var a, b int

	fmt.Print("Masukkan jumlah kelopak (a): ")
	fmt.Scan(&a)
	fmt.Print("Masukkan jumlah kelopak (b): ")
	fmt.Scan(&b)

	fmt.Printf("Bunga Sempurna antara %d dan %d: ", a, b)

	found := false
	for i := a; i <= b; i++ {
		if isPerfect(i) {
			fmt.Print(i, " ")
			found = true
		}
	}

	if !found {
		fmt.Println("Tidak ada bunga sempurna dalam rentang ini.")
	} else {
		fmt.Println()
	}
}
