package main

import (
	"flag"
	"fmt"
	"math"
	"sort"
	"time"
)

// Soal 1
func addPhones(slicePhone *[]string, phoneToAdd string) {
	*slicePhone = append(*slicePhone, phoneToAdd)
}

// Soal 2
func luasLingkaran(jariJari float64) float64 {
	return math.Pi * jariJari * jariJari
}

func kelilingLingkaran(jariJari float64) float64 {
	return 2 * math.Pi * jariJari
}

func luasPersegiPanjang(panjang, lebar int) int {
	return panjang * lebar
}

func kelilingPersegiPanjang(panjang, lebar int) int {
	return 2 * (panjang + lebar)
}

func main() {
	// Soal 1
	var phones = []string{}

	addPhones(&phones, "Xiaomi")
	addPhones(&phones, "Asus")
	addPhones(&phones, "IPhone")
	addPhones(&phones, "Samsung")
	addPhones(&phones, "Oppo")
	addPhones(&phones, "Realme")
	addPhones(&phones, "Vivo")

	// Sorting the phones
	sort.Slice(phones, func(i, j int) bool {
		return phones[i] < phones[j]
	})

	// Soal 1
	fmt.Println("-------------------------Soal 1")

	for index, phone := range phones {
		time.Sleep(time.Second * 1)
		fmt.Printf("%d. %s\n", index+1, phone)
	}

	// Soal
	fmt.Println()
	fmt.Println("-------------------------Soal 2")
	fmt.Println("Jari jari 7")
	fmt.Printf("%.3f\n", luasLingkaran(7))
	fmt.Printf("%.3f\n", kelilingLingkaran(7))
	fmt.Println("Jari jari 10")
	fmt.Printf("%.3f\n", luasLingkaran(10))
	fmt.Printf("%.3f\n", kelilingLingkaran(10))
	fmt.Println("Jari jari 15")
	fmt.Printf("%.3f\n", luasLingkaran(15))
	fmt.Printf("%.3f\n", kelilingLingkaran(15))

	// Soal 3
	fmt.Println()
	fmt.Println("-------------------------Soal 3")

	panjang := flag.Int("panjang", 9, "Masukkan panjang persegi")
	lebar := flag.Int("lebar", 5, "Masukkan panjang lebar")

	flag.Parse()

	fmt.Printf("Menghitung luas dan keliling dari persegi panjang dengan panjang %d dan lebar %d\n", *panjang, *lebar)
	fmt.Println("Luas:", luasPersegiPanjang(*panjang, *lebar))
	fmt.Println("Keliling:", kelilingPersegiPanjang(*panjang, *lebar))
}
