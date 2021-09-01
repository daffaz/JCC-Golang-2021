package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Soal 1
	var kelilingPersegiPanjang, luasSegitiga int

	var panjangPersegiPanjang string = "8"
	var lebarPersegiPanjang string = "5"

	var alasSegitiga string = "6"
	var tinggiSegitiga string = "7"

	// Init Persegi
	panjangPPInt, _ := strconv.Atoi(panjangPersegiPanjang)
	lebarPPInt, _ := strconv.Atoi(lebarPersegiPanjang)

	// Init Segitiga
	alasSgtInt, _ := strconv.Atoi(alasSegitiga)
	tinggiSgtInt, _ := strconv.Atoi(tinggiSegitiga)

	// Jawaban
	fmt.Println("================================== Soal 1")
	kelilingPersegiPanjang = (panjangPPInt + lebarPPInt) * 2
	fmt.Println("Keliling Persegi Panjang = " + strconv.Itoa(kelilingPersegiPanjang))

	luasSegitiga = (alasSgtInt * tinggiSgtInt) / 2
	fmt.Println("Luas Segitiga =", luasSegitiga)

	// Soal 2

	var nilaiJohn = 55
	var nilaiDoe = 50
	fmt.Println("\n================================== Soal 2")

	// Nilai Index John
	if nilaiJohn >= 80 {
		fmt.Println("Nilai John A")
	} else if nilaiJohn >= 70 {
		fmt.Println("Nilai John B")
	} else if nilaiJohn >= 60 {
		fmt.Println("Nilai John C")
	} else if nilaiJohn >= 50 {
		fmt.Println("Nilai John D")
	} else {
		fmt.Println("Nilai John E")
	}

	// Nilai Index Doe
	if nilaiDoe >= 80 {
		fmt.Println("Nilai Doe A")
	} else if nilaiDoe >= 70 {
		fmt.Println("Nilai Doe B")
	} else if nilaiDoe >= 60 {
		fmt.Println("Nilai Doe C")
	} else if nilaiDoe >= 50 {
		fmt.Println("Nilai Doe D")
	} else {
		fmt.Println("Nilai Doe E")
	}

	// Soal 3
	fmt.Println("\n================================== Soal 3")
	var tanggal = 27
	var bulan = 4
	var tahun = 2000

	switch bulan {
	case 1:
		fmt.Println(tanggal, "Januari", tahun)
	case 2:
		fmt.Println(tanggal, "Februari", tahun)
	case 3:
		fmt.Println(tanggal, "Maret", tahun)
	case 4:
		fmt.Println(tanggal, "April", tahun)
	case 5:
		fmt.Println(tanggal, "Mei", tahun)
	case 6:
		fmt.Println(tanggal, "Juni", tahun)
	case 7:
		fmt.Println(tanggal, "Juli", tahun)
	case 8:
		fmt.Println(tanggal, "Agustus", tahun)
	case 9:
		fmt.Println(tanggal, "September", tahun)
	case 10:
		fmt.Println(tanggal, "Oktober", tahun)
	case 11:
		fmt.Println(tanggal, "November", tahun)
	case 12:
		fmt.Println(tanggal, "Desember", tahun)
	default:
		fmt.Println("Format tanggal tidak sesuai")
	}

	// Soal 4
	fmt.Println("\n================================== Soal 4")

	switch {
	case tahun >= 1944 && tahun <= 1964:
		fmt.Println("Anda Baby boomer")
	case tahun >= 1965 && tahun <= 1979:
		fmt.Println("Anda Generasi X")
	case tahun >= 1980 && tahun <= 1994:
		fmt.Println("Anda Generasi Y (Millenials)")
	case tahun >= 1995 && tahun <= 2015:
		fmt.Println("Anda Generasi Z")
	default:
		fmt.Println("Gen Halilintar")
	}
}
