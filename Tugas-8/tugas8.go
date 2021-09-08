package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Deklarasi Pi

const PI = 3.14159265359

// Struct
// Soal 1
type segitigaSamaSisi struct {
	alas, tinggi int
}

type persegiPanjang struct {
	panjang, lebar int
}

type tabung struct {
	jariJari, tinggi float64
}

type balok struct {
	panjang, lebar, tinggi int
}

// Soal 2

type phone struct {
	name, brand string
	year        int
	colors      []string
}

// Interface
type hitungBangunDatar interface {
	luas() int
	keliling() int
}

type hitungBangunRuang interface {
	volume() float64
	luasPermukaan() float64
}

// Soal 2
type tampilDataPhone interface {
	tampilData() string
}

// Function

// Segitiga
func (s *segitigaSamaSisi) luas() int {
	return s.alas * s.tinggi / 2
}

func (s *segitigaSamaSisi) keliling() int {
	return s.alas * 3
}

// // Persegi panjang
func (p *persegiPanjang) luas() int {
	return p.panjang * p.lebar
}

func (p *persegiPanjang) keliling() int {
	return (p.panjang + p.lebar) * 2
}

// Tabung
func (t *tabung) volume() float64 {
	return PI * t.jariJari * t.jariJari * t.tinggi
}

func (t *tabung) luasPermukaan() float64 {
	return 2 * PI * t.jariJari * (t.jariJari + t.tinggi)
}

// Balok
func (b *balok) volume() float64 {
	return float64(b.panjang * b.lebar * b.tinggi)
}

func (b *balok) luasPermukaan() float64 {
	return float64(2*(b.panjang*b.lebar) + 2*(b.panjang*b.tinggi) + 2*(b.lebar*b.tinggi))
}

// Soal 2

func (p *phone) tampilData() string {
	result := fmt.Sprintf("name : %s\nbrand : %s\nyear : %d\ncolors : %s", p.name, p.brand, p.year, strings.Join(p.colors, ", "))
	return result
}

// Soal 3

func luasPersegi(sisi int, status bool) interface{} {
	switch {
	case sisi > 1:
		if status {
			return sisi * sisi
		} else {
			return sisi
		}
	default:
		if status {
			return "Maaf anda belum menginput sisi dari persegi"
		} else {
			return nil
		}
	}
}
func main() {
	// Soal 1
	fmt.Println("-------------Soal 1")

	// Deklarasi bangun data
	var segitigaSamaSisi1 hitungBangunDatar

	segitigaSamaSisi1 = &segitigaSamaSisi{
		alas:   10,
		tinggi: 20,
	}

	var persegiPanjang1 hitungBangunDatar = &persegiPanjang{
		panjang: 4,
		lebar:   8,
	}

	// Deklarasi bangun ruang
	var tabung1 hitungBangunRuang = &tabung{jariJari: 12, tinggi: 7}
	var balok1 hitungBangunRuang = &balok{panjang: 12, lebar: 5, tinggi: 6}

	// Printing method

	// Bangun datar
	fmt.Printf("Segitiga dengan alas %d dan tinggi %d\n", segitigaSamaSisi1.(*segitigaSamaSisi).alas, segitigaSamaSisi1.(*segitigaSamaSisi).tinggi)
	fmt.Printf("Luas = %d\n", segitigaSamaSisi1.luas())
	fmt.Printf("Keliling = %d\n", segitigaSamaSisi1.keliling())
	fmt.Println()

	fmt.Printf("Persegi panjang dengan panjang %d dan lebar %d\n", persegiPanjang1.(*persegiPanjang).panjang, persegiPanjang1.(*persegiPanjang).lebar)
	fmt.Printf("Luas = %d\n", persegiPanjang1.luas())
	fmt.Printf("Keliling = %d\n", persegiPanjang1.keliling())
	fmt.Println()

	// Bangun ruang

	fmt.Printf("Tabung dengan jari-jari %.f dan tinggi %.f\n", tabung1.(*tabung).jariJari, tabung1.(*tabung).tinggi)
	fmt.Printf("Volume = %.2f\n", tabung1.volume())
	fmt.Printf("Luas permukaan = %.2f\n", tabung1.luasPermukaan())
	fmt.Println()

	fmt.Printf("Balok dengan panjang %d, lebar %d dan tinggi %d\n", balok1.(*balok).panjang, balok1.(*balok).lebar, balok1.(*balok).tinggi)
	fmt.Printf("Volume = %.f\n", balok1.volume())
	fmt.Printf("Luas permukaan = %.f", balok1.luasPermukaan())
	fmt.Println()

	// Soal 2
	fmt.Println()
	fmt.Println("-------------Soal 2")
	var samsung phone = phone{
		name:   "Samsung Galaxy Note 20",
		brand:  "Samsung Galaxy Note 20",
		year:   2020,
		colors: []string{"Mystic Bronze", "Mystic White", "Mystic Black"},
	}

	fmt.Println(samsung.tampilData())

	// Soal 3
	fmt.Println()
	fmt.Println("-------------Soal 3")

	fmt.Println(luasPersegi(4, true))
	fmt.Println(luasPersegi(8, false))
	fmt.Println(luasPersegi(0, true))
	fmt.Println(luasPersegi(0, false))

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
