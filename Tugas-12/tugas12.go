package main

import (
	"fmt"
	"math"
	"sort"
	"sync"
	"time"
)

// Soal 1

func soal1(handphone []string, wg *sync.WaitGroup) {
	for i, hp := range handphone {
		wg.Add(1)
		time.Sleep(500 * time.Millisecond)
		fmt.Printf("%d. %s\n", i+1, hp)
		wg.Done()
	}
}

// Soal 2
func getMovies(ch chan string, movies ...string) {
	for _, movie := range movies {
		ch <- movie
	}
	close(ch)
}

// Soal 3
func luasLingkaran(jariJari float64, ch chan float64) {
	result := math.Pi * jariJari * jariJari
	ch <- result
}
func kelLingkaran(jariJari float64, ch chan float64) {
	result := 2 * math.Pi * jariJari
	ch <- result
}
func volumTabung(jariJari, tinggi float64, ch chan float64) {
	result := math.Pi * jariJari * jariJari * tinggi
	ch <- result
}

// Soal 4
func luasPP(panjang, lebar int, ch chan int) {
	result := panjang * lebar
	ch <- result
}
func kelilingPP(panjang, lebar int, ch chan int) {
	result := 2 * (panjang + lebar)
	ch <- result
}
func volumBalok(panjang, lebar, tinggi int, ch chan int) {
	result := panjang * lebar * tinggi
	ch <- result
}

func main() {
	// Soal 1
	var phones = []string{"Xiaomi", "Asus", "Iphone", "Samsung", "Oppo", "Realme", "Vivo"}

	sort.Slice(phones, func(i, j int) bool {
		return phones[i] < phones[j]
	})
	var wg sync.WaitGroup

	soal1(phones, &wg)
	wg.Wait()
	// Soal 2

	fmt.Println()

	var movies = []string{"Harry Potter", "LOTR", "SpiderMan", "Logan", "Avengers", "Insidious", "Toy Story"}

	moviesChannel := make(chan string)

	go getMovies(moviesChannel, movies...)

	var nomor = 1
	for value := range moviesChannel {
		fmt.Printf("%d. %s\n", nomor, value)
		nomor++
	}

	// Soal 3
	fmt.Println()

	soal3Ch1 := make(chan float64)
	soal3Ch2 := make(chan float64)
	soal3Ch3 := make(chan float64)
	go luasLingkaran(8, soal3Ch1)
	go kelLingkaran(14, soal3Ch2)
	go volumTabung(20, 10, soal3Ch3)

	luasLingkar := <-soal3Ch1
	fmt.Println("Luas lingkaran:", luasLingkar)

	kelilingLingkar := <-soal3Ch2
	fmt.Println("Keliling lingkaran:", kelilingLingkar)

	volumTab := <-soal3Ch3
	fmt.Println("Volum tabung:", volumTab)

	// Soal 4

	fmt.Println()

	soal4Ch1 := make(chan int)
	soal4Ch2 := make(chan int)
	soal4Ch3 := make(chan int)

	go luasPP(10, 5, soal4Ch1)
	go kelilingPP(10, 5, soal4Ch2)
	go volumBalok(10, 5, 8, soal4Ch3)

	for i := 0; i < 3; i++ {
		select {
		case luasPerpanjang := <-soal4Ch1:
			fmt.Println("Luas Persegi panjang:", luasPerpanjang)
		case kelilingPerpanjang := <-soal4Ch2:
			fmt.Println("Keliling Persegi panjang:", kelilingPerpanjang)
		case volBalok := <-soal4Ch3:
			fmt.Println("Volum Balok:", volBalok)
		}
	}
}
