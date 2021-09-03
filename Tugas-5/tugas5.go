package main

import (
	"fmt"
	"strings"
)

// Soal 1
func luasPersegiPanjang(panjang, lebar uint32) uint32 {
	return panjang * lebar
}

func kelilingPersegiPanjang(panjang, lebar uint32) uint32 {
	return (panjang + lebar) * 2
}

func volumeBalok(panjang, lebar, tinggi uint32) uint32 {
	return panjang * lebar * tinggi
}

// Soal 2
func introduce(name, gender, job, umur string) (result string) {
	if gender == "perempuan" {
		result = "Bu " + name + " adalah seorang " + job + " yang berusia " + umur + " tahun"
	} else {
		result = "Pak " + name + " adalah seorang " + job + " yang berusia " + umur + " tahun"
	}
	return
}

// Soal 3
func buahFavorit(nama string, sliceBuah ...string) string {
	result := "Halo nama saya " + nama + " dan buah favorit saya adalah \"" + strings.Join(sliceBuah, ", ") + "\""
	return result
}

func main() {
	// Soal 1
	luas := luasPersegiPanjang(3, 4)
	keliling := kelilingPersegiPanjang(3, 4)
	volume := volumeBalok(3, 4, 5)

	fmt.Println("+++++++++++++ Nomor 1")
	fmt.Println("Luas:", luas)
	fmt.Println("Keliling:", keliling)
	fmt.Println("Volume:", volume)

	// Soal 2

	fmt.Println("\n+++++++++++++ Nomor 2")

	john := introduce("John", "laki-laki", "penulis", "30")
	fmt.Println(john) // Menampilkan "Pak John adalah seorang penulis yang berusia 30 tahun"

	sarah := introduce("Sarah", "perempuan", "model", "28")
	fmt.Println(sarah) // Menampilkan "Bu Sarah adalah seorang model yang berusia 28 tahun"

	// Soal 3
	fmt.Println("\n+++++++++++++ Nomor 3")

	var buah = []string{"semangka", "jeruk", "melon", "pepaya"}
	var buahFavoritJohn = buahFavorit("John", buah...)

	fmt.Println(buahFavoritJohn)

	// Soal 4
	fmt.Println("\n+++++++++++++ Nomor 4")

	var dataFilm = []map[string]string{}

	var tambahDataFilm = func(title, jam, genre, tahun string) {
		dataFilm = append(dataFilm, map[string]string{"genre": genre, "jam": jam, "tahun": tahun, "title": title})
	}

	tambahDataFilm("LOTR", "2 jam", "action", "1999")
	tambahDataFilm("avenger", "2 jam", "action", "2019")
	tambahDataFilm("spiderman", "2 jam", "action", "2004")
	tambahDataFilm("juon", "2 jam", "horror", "2004")

	for _, item := range dataFilm {
		fmt.Println(item)
	}

}
