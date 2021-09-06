package main

import "fmt"

// Soal 1

const PI float64 = 3.14159265359

func cariLuasLingkaran(reference *float64, r float64) {
	*reference = PI * r * r
}

func cariKelilingLingkaran(reference *float64, r float64) {
	*reference = 2 * PI * r
}

// Soal 2

func introduce(reference *string, name, gender, job, age string) {
	var genderSentence string
	if gender == "perempuan" {
		genderSentence = "Bu"
	} else {
		genderSentence = "Pak"
	}

	*reference = genderSentence + " " + name + " adalah seorang " + job + " yang berusia " + age + " tahun"
}

// Soal 3

func tambahBuah(reference *[]string, buah string) {
	*reference = append(*reference, buah)
}

// Soal 4

func tambahDataFilm(title, duration, genre, year string, reference *[]map[string]string) {
	*reference = append(*reference, map[string]string{"title": title, "duration": duration, "genre": genre, "year": year})
}
func main() {
	// Soal 1

	var luasLigkaran float64
	var kelilingLingkaran float64

	cariLuasLingkaran(&luasLigkaran, 14)
	cariKelilingLingkaran(&kelilingLingkaran, 14)

	fmt.Printf("Luas lingkaran: %.2f\n", luasLigkaran)
	fmt.Printf("Luas lingkaran: %.2f\n", kelilingLingkaran)

	// Soal 2
	fmt.Println()
	var sentence string
	introduce(&sentence, "John", "laki-laki", "penulis", "30")

	fmt.Println(sentence) // Menampilkan "Pak John adalah seorang penulis yang berusia 30 tahun"
	introduce(&sentence, "Sarah", "perempuan", "model", "28")

	fmt.Println(sentence) // Menampilkan "Bu Sarah adalah seorang model yang berusia 28 tahun"

	// Soal 3
	fmt.Println()
	var buah []string = []string{}

	// Proses tambah buah
	tambahBuah(&buah, "Jeruk")
	tambahBuah(&buah, "Semangka")
	tambahBuah(&buah, "Mangga")
	tambahBuah(&buah, "Strawberry")
	tambahBuah(&buah, "Durian")
	tambahBuah(&buah, "Manggis")
	tambahBuah(&buah, "Alpukat")

	for i, namaBuah := range buah {
		fmt.Printf("%d. %s\n", i+1, namaBuah)
	}

	// Soal 4

	fmt.Println()
	var dataFilm = []map[string]string{}

	tambahDataFilm("LOTR", "2 jam", "action", "1999", &dataFilm)
	tambahDataFilm("avenger", "2 jam", "action", "2019", &dataFilm)
	tambahDataFilm("spiderman", "2 jam", "action", "2004", &dataFilm)
	tambahDataFilm("juon", "2 jam", "horror", "2004", &dataFilm)

	for index, film := range dataFilm {
		fmt.Printf("%d. title: %s\n   duration: %s\n   genre: %s\n   year: %s\n\n", index+1, film["title"], film["duration"], film["genre"], film["year"])
	}
}
