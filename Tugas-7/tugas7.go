package main

import (
	"fmt"
)

type buah struct {
	Nama       string
	Warna      string
	AdaBijinya bool
	Harga      int
}

// Function tambahan untuk soal 1

func adaBiji(buah bool) string {
	var berbiji string

	if buah {
		berbiji = "Ada"
	} else {
		berbiji = "Tidak"
	}

	return berbiji
}

// Soal 2
type segitiga struct {
	alas, tinggi int
}

type persegi struct {
	sisi int
}

type persegiPanjang struct {
	panjang, lebar int
}

func (s segitiga) luasSegitiga() int {
	return s.alas * s.tinggi / 2
}

func (p persegi) luasPersegi() int {
	return p.sisi * p.sisi
}

func (pp persegiPanjang) luasPersegiPanjang() int {
	return pp.panjang * pp.lebar
}

// Soal 3
type phone struct {
	name, brand string
	year        int
	colors      []string
}

func (p *phone) addColors(color string) {
	p.colors = append(p.colors, color)
}

// Soal 4
type movie struct {
	title, genre   string
	duration, year int
}

func tambahDataFilm(title string, duration int, genre string, year int, reference *[]movie) {
	durationInHour := duration / 60
	*reference = append(*reference, movie{title: title, duration: durationInHour, genre: genre, year: year})
}

func main() {
	// Soal 1
	var Nanas buah = buah{}

	Nanas.Nama = "Nanas"
	Nanas.Warna = "Kuning"
	Nanas.AdaBijinya = false
	Nanas.Harga = 9000

	var Jeruk buah = buah{"Jeruk", "Oranye", true, 8000}

	var Semangka buah = buah{
		Nama:       "Semangka",
		Warna:      "Hijau & Merah",
		AdaBijinya: true,
		Harga:      10000,
	}

	Pisang := buah{
		Nama:       "Pisang",
		Warna:      "Kuning",
		AdaBijinya: false,
		Harga:      5000,
	}

	fmt.Printf("Nama \t\t Warna \t\t Ada Bijinya \t Harga\n")
	fmt.Println("-------------------------------------------------------")
	fmt.Printf("%s \t\t %s \t %s \t\t %d\n", Nanas.Nama, Nanas.Warna, adaBiji(Nanas.AdaBijinya), Nanas.Harga)
	fmt.Printf("%s \t\t %s \t %s \t\t %d\n", Jeruk.Nama, Jeruk.Warna, adaBiji(Jeruk.AdaBijinya), Jeruk.Harga)
	fmt.Printf("%s \t %s \t %s \t\t %d\n", Semangka.Nama, Semangka.Warna, adaBiji(Semangka.AdaBijinya), Semangka.Harga)
	fmt.Printf("%s \t\t %s \t %s \t\t %d\n", Pisang.Nama, Pisang.Warna, adaBiji(Pisang.AdaBijinya), Pisang.Harga)

	// Soal 2

	fmt.Println()
	// Segitiga
	var segitigaSamaSisi segitiga = segitiga{alas: 12, tinggi: 12}
	luasSegitigaVar := segitigaSamaSisi.luasSegitiga()
	fmt.Printf("Luas segitiga dengan alas %d dan tinggi %d = %d\n", segitigaSamaSisi.alas, segitigaSamaSisi.tinggi, luasSegitigaVar)

	// Persegi
	var persegiVar persegi = persegi{sisi: 10}
	luasPersegiVar := persegiVar.luasPersegi()
	fmt.Printf("Luas persegi dengan sisi %d = %d\n", persegiVar.sisi, luasPersegiVar)

	var persegiPanjangVar persegiPanjang = persegiPanjang{panjang: 5, lebar: 5}
	luasPersegiPanjangVar := persegiPanjangVar.luasPersegiPanjang()
	fmt.Printf("Luas persegi panjang dengan panjang %d dan lebar %d = %d\n", persegiPanjangVar.panjang, persegiPanjangVar.lebar, luasPersegiPanjangVar)

	// Soal 3

	fmt.Println()
	iPhone := phone{name: "IPhone", brand: "Apple", year: 2016}

	iPhone.addColors("Red")
	iPhone.addColors("Black")
	iPhone.addColors("Blue")

	fmt.Println(iPhone)
	fmt.Println("Warna pada iPhone")
	for i, warna := range iPhone.colors {
		fmt.Printf("%d. %s\n", i+1, warna)
	}

	// Soal 4

	fmt.Println()

	var dataFilm = []movie{}

	tambahDataFilm("LOTR", 120, "action", 1999, &dataFilm)
	tambahDataFilm("avenger", 120, "action", 2019, &dataFilm)
	tambahDataFilm("spiderman", 120, "action", 2004, &dataFilm)
	tambahDataFilm("juon", 120, "horror", 2004, &dataFilm)

	// fmt.Println(dataFilm)
	for i, val := range dataFilm {
		fmt.Printf("%d. title: %s\n   duration: %d jam\n   genre: %s\n   year: %d\n\n", i+1, val.title, val.duration, val.genre, val.year)
	}
}
