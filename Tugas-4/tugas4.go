package main

import (
	"fmt"
)

func main() {
	// Soal 1
	fmt.Println("-------------------- SOAL 1")
	for i := 1; i <= 20; i++ {
		if i%2 == 1 {
			if i%3 == 0 {
				fmt.Printf("%d-  I Love Coding\n", i)
			} else {
				fmt.Printf("%d - JCC\n", i)
			}
		} else {
			fmt.Printf("%d - Candradimuka\n", i)
		}
	}
	fmt.Println("--------------------")

	// Soal 2
	fmt.Println("-------------------- SOAL 2")

	// Lakukan inisialisai variabel board, untuk menampung hashtag
	var board string

	// Setiap iterasi, tambahkan hashtag sebanyak satu lalu print variabel board
	for i := 0; i < 7; i++ {
		board += "#"
		fmt.Println(board)
	}

	fmt.Println("--------------------")

	// Soal 3
	fmt.Println("-------------------- SOAL 3")

	var kalimat = [...]string{"aku", "dan", "saya", "sangat", "senang", "belajar", "golang"}

	var newKalimat = kalimat[2:]
	fmt.Println(newKalimat)

	fmt.Println("--------------------")

	// Soal 4
	fmt.Println("-------------------- SOAL 4")
	var sayuran = []string{}
	sayuran = []string{
		"Bayam",
		"Buncis",
		"Kangkung",
		"Kubis",
		"Seledri",
		"Tauge",
		"Timun",
	}

	// Solusi kedua bisa seperti ini
	// var sayuran = make([]string, 7)
	// sayuran[0] = "Bayam"
	// sayuran[1] = "Buncis"
	// sayuran[2] = "Kangkung"
	// sayuran[3] = "Kubis"
	// sayuran[4] = "Seledri"
	// sayuran[5] = "Tauge"
	// sayuran[6] = "Timun"

	for index, sayur := range sayuran {
		// fmt.Printf("%d. %s \n", index, sayur)
		fmt.Printf("%d. %s \n", index+1, sayur)
	}
	fmt.Println("--------------------")

	// Soal 5
	fmt.Println("-------------------- SOAL 5")

	var satuan = map[string]int{
		"panjang": 7,
		"lebar":   4,
		"tinggi":  6,
	}

	var volumeBalok int = 1
	for key, value := range satuan {
		fmt.Printf("%s = %d\n", key, value)
		volumeBalok *= value
	}
	fmt.Printf("Volume Balok = %d\n", volumeBalok)

	fmt.Println("--------------------")
}
