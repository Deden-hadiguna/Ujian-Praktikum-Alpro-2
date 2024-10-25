//2311102286
//Deden Hadiguna

package main

import "fmt"

func main() {
	var b int
	fmt.Print("Masukkan jumlah tiket: ")
	fmt.Scan(&b)

	fmt.Printf("Masukkan jenis kursi (biasa/VIP): %d\n", b)

	for i := 1; i <= b; i++ {
		if b%i == 0 {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()

	if isPrime(b) {
		fmt.Println("Apaakah anda member?: true")
	} else {
		fmt.Println("Apakah anda member?: false")
	}
}

func isPrime(n int) bool {
	if n <= 2 {
		return false
	}
	for i := 3; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
