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

func kelilingSegitigaSamaSisi(sisi int, hitungKeliling bool) (string, error) {
	if hitungKeliling == true {
		if sisi > 0 {
			return "keliling segitiga sama sisinya dengan sisi " + strconv.Itoa(sisi) + " cm adalah " + strconv.Itoa(sisi*3) + " cm", nil
		} else {
			return "", errors.New("Maaf anda belum menginput sisi dari segitiga sama sisi")
		}
	} else {
		if sisi > 0 {
			return strconv.Itoa(sisi), nil
		} else {
			defer recover()
			panic("wow")
			return "", errors.New("Maaf anda belum menginput sisi dari segitiga sama sisi")
		}
	}
}

func main() {
	// Soal 1
	defer printJabarcoding("Candradimuka Jabar Coding Camp", 2021)

	// Soal 2
	fmt.Println()
	fmt.Println("-------------Soal 2")

	fmt.Println(kelilingSegitigaSamaSisi(4, true))
	fmt.Println(kelilingSegitigaSamaSisi(8, false))
	fmt.Println(kelilingSegitigaSamaSisi(0, true))
	fmt.Println(kelilingSegitigaSamaSisi(0, false))
}
