package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// Soal 2

func hitungLuasSegitiga(alas, tinggi int) int {
	return alas * tinggi / 2
}

func hitungKelilingSegitiga(alas int) int {
	return alas * 3
}

func hitungLuasPersegi(sisi int) int {
	return sisi * sisi
}

func hitungKelilingPersegi(sisi int) int {
	return sisi * 4
}

func hitungLuasPP(panjang, lebar int) int {
	return panjang * lebar
}

func hitungKelilingPP(panjang, lebar int) int {
	return 2 * (panjang + lebar)
}

func hitungLuasLingkaran(jariJari float64) float64 {
	return math.Pi * jariJari * jariJari
}

func hitungKelilingLingkaran(jariJari float64) float64 {
	return 2 * math.Pi * jariJari
}

func hitungLuasJajarGenjang(alas, tinggi int) int {
	return alas * tinggi
}

func hitungKelilingJajarGenjang(alas, sisi int) int {
	return 2 * (alas + sisi)
}
func BangunDatarHandler() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		paramHitung := r.URL.Query().Get("hitung")
		if params.ByName("bangundatar") == "segitiga-sama-sisi" {
			if paramHitung == "luas" {
				alas, _ := strconv.Atoi(r.URL.Query().Get("alas"))
				tinggi, _ := strconv.Atoi(r.URL.Query().Get("tinggi"))
				luas := hitungLuasSegitiga(alas, tinggi)
				fmt.Fprintf(w, "Luas dari segitiga dengan alas %d dan tinggi %d adalah %d", alas, tinggi, luas)
			} else if paramHitung == "keliling" {
				alas, _ := strconv.Atoi(r.URL.Query().Get("alas"))
				keliling := hitungKelilingSegitiga(alas)
				fmt.Fprintf(w, "Keliling dari segitiga dengan alas %d  adalah %d", alas, keliling)
			} else {
				fmt.Fprintf(w, "Hmmm sepertinya parameter yang anda masukkan salah")
			}
		} else if params.ByName("bangundatar") == "persegi" {
			if paramHitung == "luas" {
				sisi, _ := strconv.Atoi(r.URL.Query().Get("sisi"))
				luas := hitungLuasPersegi(sisi)
				fmt.Fprintf(w, "Luas dari persegi dengan sisi %d adalah %d", sisi, luas)
			} else if paramHitung == "keliling" {
				sisi, _ := strconv.Atoi(r.URL.Query().Get("sisi"))
				keliling := hitungKelilingPersegi(sisi)
				fmt.Fprintf(w, "Keliling dari persegi dengan sisi %d adalah %d", sisi, keliling)
			} else {
				fmt.Fprintf(w, "Hmmm sepertinya parameter yang anda masukkan salah")
			}
		} else if params.ByName("bangundatar") == "persegi-panjang" {
			if paramHitung == "luas" {
				panjang, _ := strconv.Atoi(r.URL.Query().Get("panjang"))
				lebar, _ := strconv.Atoi(r.URL.Query().Get("lebar"))
				luas := hitungLuasPP(panjang, lebar)
				fmt.Fprintf(w, "Luas dari persegi panjang dengan panjang %d dan lebar %d adalah %d", panjang, lebar, luas)
			} else if paramHitung == "keliling" {
				panjang, _ := strconv.Atoi(r.URL.Query().Get("panjang"))
				lebar, _ := strconv.Atoi(r.URL.Query().Get("lebar"))
				keliling := hitungKelilingPP(panjang, lebar)
				fmt.Fprintf(w, "Keliling dari persegi panjang dengan panjang %d dan lebar %d adalah %d", panjang, lebar, keliling)
			} else {
				fmt.Fprintf(w, "Hmmm sepertinya parameter yang anda masukkan salah")
			}
		} else if params.ByName("bangundatar") == "lingkaran" {
			if paramHitung == "luas" {
				jariJari := r.URL.Query().Get("jariJari")
				jariJariFix, _ := strconv.ParseFloat(jariJari, 64)
				luas := hitungLuasLingkaran(jariJariFix)
				fmt.Fprintf(w, "Luas dari lingkaran dengan jari-jari %v adalah %.3f", jariJariFix, luas)
			} else if paramHitung == "keliling" {
				jariJari := r.URL.Query().Get("jariJari")
				jariJariFix, _ := strconv.ParseFloat(jariJari, 64)
				keliling := hitungKelilingLingkaran(jariJariFix)
				fmt.Fprintf(w, "Keliling dari lingkaran dengan jari-jari %v adalah %.3f", jariJariFix, keliling)
			} else {
				fmt.Fprintf(w, "Hmmm sepertinya parameter yang anda masukkan salah")
			}
		} else if params.ByName("bangundatar") == "jajar-genjang" {
			if paramHitung == "luas" {
				alas, _ := strconv.Atoi(r.URL.Query().Get("alas"))
				tinggi, _ := strconv.Atoi(r.URL.Query().Get("tinggi"))
				luas := hitungLuasJajarGenjang(alas, tinggi)
				fmt.Fprintf(w, "Luas dari jajargenjang dengan alas %d dan tinggi %d adalah %d", alas, tinggi, luas)
			} else if paramHitung == "keliling" {
				alas, _ := strconv.Atoi(r.URL.Query().Get("alas"))
				sisi, _ := strconv.Atoi(r.URL.Query().Get("sisi"))
				luas := hitungKelilingJajarGenjang(alas, sisi)
				fmt.Fprintf(w, "Keliling dari jajargenjang dengan alas %d dan sisi %d adalah %d", alas, sisi, luas)
			} else {
				fmt.Fprintf(w, "Hmmm sepertinya parameter yang anda masukkan salah")
			}
		} else {
			fmt.Fprintf(w, "Bangun datar \"%s\" tidak terdaftar pada penghitungan", params.ByName("bangundatar"))
		}
	}
}

// Soal 5
func middleWare(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		username, password, isOk := r.BasicAuth()

		if !isOk {
			w.Write([]byte("Isi username dan password terlebih dahulu"))
			return
		}

		if username == "admin" && password == "password" || username == "editor" && password == "secret" || username == "trainer" && password == "rahasia" {
			next(w, r, ps)
			return
		}
		w.Write([]byte("Username atau password tidak sesuai"))
	}
}

func main() {
	router := httprouter.New()

	// Soal 2
	router.GET("/bangun-datar/:bangundatar", middleWare(BangunDatarHandler()))

	// Soal 3
	// router.GET("/books/", middleWare())

	log.Fatal(http.ListenAndServe(":10000", router))
}
