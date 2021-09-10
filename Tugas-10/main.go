package main

import (
	"fmt"
	"strconv"
	"strings"
	"tugas10/tugas8"
)

func main() {
	// Soal 1
	fmt.Println("-------------Soal 1")

	// Deklarasi bangun data
	var SegitigaSamaSisi1 tugas8.HitungBangunDatar = tugas8.SegitigaSamaSisi{
		Alas:   10,
		Tinggi: 20,
	}

	var PersegiPanjang1 tugas8.HitungBangunDatar = tugas8.PersegiPanjang{
		Panjang: 4,
		Lebar:   8,
	}

	// Deklarasi bangun ruang
	var Tabung1 tugas8.HitungBangunRuang = tugas8.Tabung{JariJari: 12, Tinggi: 7}
	var Balok1 tugas8.HitungBangunRuang = tugas8.Balok{Panjang: 12, Lebar: 5, Tinggi: 6}

	// Printing method

	// Bangun datar
	fmt.Printf("Segitiga dengan Alas %d dan Tinggi %d\n", SegitigaSamaSisi1.(tugas8.SegitigaSamaSisi).Alas, SegitigaSamaSisi1.(tugas8.SegitigaSamaSisi).Tinggi)
	fmt.Printf("Luas = %d\n", SegitigaSamaSisi1.Luas())
	fmt.Printf("Keliling = %d\n", SegitigaSamaSisi1.Keliling())
	fmt.Println()

	fmt.Printf("Persegi Panjang dengan Panjang %d dan Lebar %d\n", PersegiPanjang1.(tugas8.PersegiPanjang).Panjang, PersegiPanjang1.(tugas8.PersegiPanjang).Lebar)
	fmt.Printf("Luas = %d\n", PersegiPanjang1.Luas())
	fmt.Printf("Keliling = %d\n", PersegiPanjang1.Keliling())
	fmt.Println()

	// Bangun ruang

	fmt.Printf("Tabung dengan jari-jari %.f dan Tinggi %.f\n", Tabung1.(tugas8.Tabung).JariJari, Tabung1.(tugas8.Tabung).Tinggi)
	fmt.Printf("Volume = %.2f\n", Tabung1.Volume())
	fmt.Printf("Luas permukaan = %.2f\n", Tabung1.LuasPermukaan())
	fmt.Println()

	fmt.Printf("Balok dengan Panjang %d, Lebar %d dan Tinggi %d\n", Balok1.(tugas8.Balok).Panjang, Balok1.(tugas8.Balok).Lebar, Balok1.(tugas8.Balok).Tinggi)
	fmt.Printf("Volume = %.f\n", Balok1.Volume())
	fmt.Printf("Luas permukaan = %.f", Balok1.LuasPermukaan())
	fmt.Println()

	// Soal 2
	fmt.Println()
	fmt.Println("-------------Soal 2")
	var samsung tugas8.TampilDataPhone = tugas8.Phone{
		Name:   "Samsung Galaxy Note 20",
		Brand:  "Samsung Galaxy Note 20",
		Year:   2020,
		Colors: []string{"Mystic Bronze", "Mystic White", "Mystic Black"},
	}

	fmt.Println(samsung.TampilData())

	// Soal 3
	fmt.Println()
	fmt.Println("-------------Soal 3")

	fmt.Println(tugas8.LuasPersegi(4, true))
	fmt.Println(tugas8.LuasPersegi(8, false))
	fmt.Println(tugas8.LuasPersegi(0, true))
	fmt.Println(tugas8.LuasPersegi(0, false))

	// Soal 4
	fmt.Println()
	fmt.Println("-------------Soal 4")

	var prefix interface{} = "hasil penjumlahan dari "

	var kumpulanAngkaPertama interface{} = []int{6, 8}

	var kumpulanAngkaKedua interface{} = []int{12, 14}

	var kumpulanAngkaFixed = append(kumpulanAngkaPertama.([]int), kumpulanAngkaKedua.([]int)...)
	var resultKumpulanAngka int = 0

	var kumpulanAngkaInString = make([]string, len(kumpulanAngkaFixed))

	for i, angka := range kumpulanAngkaFixed {
		resultKumpulanAngka += angka
		kumpulanAngkaInString[i] = strconv.Itoa(angka)
	}
	fmt.Println(prefix, strings.Join(kumpulanAngkaInString, " + "), "=", strconv.Itoa(resultKumpulanAngka))
}
