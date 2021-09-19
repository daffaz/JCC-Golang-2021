package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net/http"
	"quiz3/controller"
	"quiz3/model"
	"quiz3/utils"
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
func removeUseless(s map[string]string) map[string]string {
	var r map[string]string = map[string]string{}
	for key, str := range s {
		if str != "" {
			r[key] = str
		}
	}
	return r
}

// Soal 3
func GetDataBuku() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		context, cancel := context.WithCancel(context.Background())
		defer cancel()

		paramTitle := r.URL.Query().Get("title")
		paramminYear := r.URL.Query().Get("minYear")
		parammaxYear := r.URL.Query().Get("maxYear")
		paramminPage := r.URL.Query().Get("minPage")
		parammaxPage := r.URL.Query().Get("maxPage")
		paramsort := r.URL.Query().Get("sort")

		// utils.ResponseToJSON(w, paramTitle, http.StatusBadGateway)
		// return

		allParams := map[string]string{
			"paramTitle":   paramTitle,
			"paramminYear": paramminYear,
			"parammaxYear": parammaxYear,
			"paramminPage": paramminPage,
			"parammaxPage": parammaxPage,
			"paramsort":    paramsort,
		}
		paramram := removeUseless(allParams)

		dataBuku, err := controller.GetAllBooks(context, paramram)
		if err != nil {
			utils.ResponseToJSON(w, err, http.StatusBadRequest)
			return
		}

		if dataBuku == nil {
			utils.ResponseToJSON(w, "Data tidak ditemukan", http.StatusNotFound)
			return
		}

		utils.ResponseToJSON(w, dataBuku, http.StatusOK)
	}
}

func PostBuku() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		if r.Header.Get("Content-type") != "application/json" {
			utils.ResponseToJSON(w, "Only accepting application/json POST method", http.StatusBadRequest)
			return
		}
		context, cancel := context.WithCancel(context.Background())
		defer cancel()

		var bukuDummy model.Book2
		if err := json.NewDecoder(r.Body).Decode(&bukuDummy); err != nil {
			utils.ResponseToJSON(w, err, http.StatusBadRequest)
			return
		}

		somePrice := strconv.Itoa(bukuDummy.Price)
		var buku model.Book = model.Book{
			Id:                bukuDummy.Id,
			Title:             bukuDummy.Title,
			Description:       bukuDummy.Description,
			ImageUrl:          bukuDummy.ImageUrl,
			ReleaseYear:       bukuDummy.ReleaseYear,
			Price:             somePrice,
			TotalPage:         bukuDummy.TotalPage,
			KategoriKetebalan: bukuDummy.KategoriKetebalan,
			CreatedAt:         bukuDummy.CreatedAt,
			UpdatedAt:         bukuDummy.UpdatedAt,
		}

		_, errLink := controller.ValidateUrl(buku.ImageUrl)
		_, errYear := controller.ValidateYear(buku.ReleaseYear)

		if errLink != nil && errYear != nil {
			utils.ResponseToJSON(w, "Tahun dan image url tidak valid", http.StatusBadRequest)
			return
		}
		if errYear != nil {
			utils.ResponseToJSON(w, "Tahun tidak boleh kurang dari 1980 dan tidak boleh lebih dari 2021", http.StatusBadRequest)
			return
		}
		if errLink != nil {
			utils.ResponseToJSON(w, "Image url tidak valid", http.StatusBadRequest)
			return
		}
		if err := controller.InsertBuku(context, buku); err != nil {
			return
		}

		response := map[string]interface{}{
			"message": "Post success",
		}
		utils.ResponseToJSON(w, response, http.StatusCreated)
	}
}

func UpdateDataBuku() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		if r.Header.Get("Content-type") != "application/json" {
			utils.ResponseToJSON(w, "Only accepting application/json POST method", http.StatusBadRequest)
			return
		}
		context, cancel := context.WithCancel(context.Background())
		defer cancel()

		var bukuDummy model.Book2
		if err := json.NewDecoder(r.Body).Decode(&bukuDummy); err != nil {
			utils.ResponseToJSON(w, err, http.StatusBadRequest)
			return
		}

		somePrice := strconv.Itoa(bukuDummy.Price)
		var buku model.Book = model.Book{
			Id:                bukuDummy.Id,
			Title:             bukuDummy.Title,
			Description:       bukuDummy.Description,
			ImageUrl:          bukuDummy.ImageUrl,
			ReleaseYear:       bukuDummy.ReleaseYear,
			Price:             somePrice,
			TotalPage:         bukuDummy.TotalPage,
			KategoriKetebalan: bukuDummy.KategoriKetebalan,
			CreatedAt:         bukuDummy.CreatedAt,
			UpdatedAt:         bukuDummy.UpdatedAt,
		}

		var idBuku = params.ByName("id")

		if err := controller.UpdateBuku(context, buku, idBuku); err != nil {
			utils.ResponseToJSON(w, err, http.StatusInternalServerError)
			return
		}

		response := map[string]interface{}{
			"message": "Update success",
		}
		utils.ResponseToJSON(w, response, http.StatusCreated)
	}
}

func DeleteDataBuku() httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		context, cancel := context.WithCancel(context.Background())
		defer cancel()

		var idBuku = params.ByName("id")
		if err := controller.DeleteBuku(context, idBuku); err != nil {
			utils.ResponseToJSON(w, err, http.StatusInternalServerError)
			return
		}

		res := map[string]string{
			"status": "Delete Success",
		}
		utils.ResponseToJSON(w, res, http.StatusOK)
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
	router.GET("/books/", middleWare(GetDataBuku()))
	router.POST("/books/create", middleWare(PostBuku()))
	router.PUT("/books/update/:id", middleWare(UpdateDataBuku()))
	router.DELETE("/books/delete/:id", middleWare(DeleteDataBuku()))

	log.Println("Running in port :10000")
	log.Fatal(http.ListenAndServe(":10000", router))
}
