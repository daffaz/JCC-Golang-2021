package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// Soal 1
	var word1 string = "Jabar"
	var word2 string = "Coding"
	var word3 string = "Camp"
	var word4 string = "Golang"
	var word5 string = "2021"

	fmt.Println("========= SOAL 1 ==========")
	fmt.Println(word1, word2, word3, word4, word5)
	fmt.Println("===========================")

	// Soal 2
	halo := "Halo Dunia"
	halo = strings.Replace(halo, "Dunia", "Golang", 1)

	fmt.Println("\n========= SOAL 2 ==========")
	fmt.Println(halo)
	fmt.Println("===========================")

	// Soal 3
	var kataPertama = "saya"
	var kataKedua = "senang"
	var kataKetiga = "belajar"
	var kataKeempat = "golang"

	var outputSoal3 string = kataPertama + " " + strings.Replace(kataKedua, "s", "S", 1) + " " + strings.Replace(kataKetiga, "r", "R", 1) + " " + strings.ToUpper(kataKeempat)
	fmt.Println("\n========= SOAL 3 ==========")
	fmt.Println(outputSoal3)
	fmt.Println("===========================")

	// Soal 4

	var angkaPertama = "8"
	var angkaKedua = "5"
	var angkaKetiga = "6"
	var angkaKeempat = "7"

	var angka1, _ = strconv.Atoi(angkaPertama)
	var angka2, _ = strconv.Atoi(angkaKedua)
	var angka3, _ = strconv.Atoi(angkaKetiga)
	var angka4, _ = strconv.Atoi(angkaKeempat)

	fmt.Println("\n========= SOAL 4 ==========")
	fmt.Println(angka1 + angka2 + angka3 + angka4)
	fmt.Println("===========================")

	// Soal 5

	kalimat := "halo halo bandung"
	angka := 2021

	kalimat = strings.Replace(kalimat, "halo", "Hi", -1)

	fmt.Println("\n========= SOAL 5 ==========")
	fmt.Printf("%q - %d\n", kalimat, angka)
	fmt.Println("===========================")
}
