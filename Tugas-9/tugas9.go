package main

import (
	"errors"
	"fmt"
	"strconv"
)

func printJabarcoding(kalimat string, tahun int) {
	fmt.Println("-------------Soal 1")
	fmt.Println(kalimat, tahun)
}

func kelilingSegitigaSamaSisi(sisi int, hitungKeliling bool) (str string, err error) {
	if hitungKeliling == true {
		if sisi > 0 {
			str = "keliling segitiga sama sisinya dengan sisi " + strconv.Itoa(sisi) + " cm adalah " + strconv.Itoa(sisi*3) + " cm"
			err = nil
		} else {
			str = ""
			err = errors.New("Maaf anda belum menginput sisi dari segitiga sama sisi")
		}
	} else {
		if sisi > 0 {
			str = strconv.Itoa(sisi)
			err = nil
		} else {
			defer func() {
				if r := recover(); r != nil {
					str = "Maaf anda belum menginput sisi dari segitiga sama sisi"
					err = errors.New(r.(string))
				}
			}()
			panic("PANIC!!")
		}
	}
	return
}

// Soal 3

func tambahAngka(angka int, varAngka *int) {
	*varAngka += angka
}

func deferSoal3(totalAngka *int) {
	fmt.Println(*totalAngka)
	fmt.Println("Soal 3 selesai")
}

func main() {
	angka := 1
	defer deferSoal3(&angka)
	// Soal 1
	defer printJabarcoding("Candradimuka Jabar Coding Camp", 2021)
	// End soal 1

	// Soal 2
	fmt.Println()
	fmt.Println("-------------Soal 2")

	fmt.Println(kelilingSegitigaSamaSisi(4, true))
	fmt.Println(kelilingSegitigaSamaSisi(8, false))
	fmt.Println(kelilingSegitigaSamaSisi(0, true))
	fmt.Println(kelilingSegitigaSamaSisi(0, false))

	// Soal 3
	fmt.Println()
	fmt.Println("-------------Soal 3")

	// deklarasi variabel angka ini simpan di baris pertama func main
	// defer deferSoal3(&angka)
	tambahAngka(7, &angka)
	tambahAngka(6, &angka)
	tambahAngka(-1, &angka)
	tambahAngka(9, &angka)
}
