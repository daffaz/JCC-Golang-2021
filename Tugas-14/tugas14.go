package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

// Soal 1
type NilaiMahasiswa struct {
	Nama, MataKuliah, IndeksNilai string
	Nilai, ID                     uint
}

func getIndexNilai(nilai int) string {
	switch {
	case nilai >= 80:
		return "A"
	case nilai >= 70:
		return "B"
	case nilai >= 60:
		return "C"
	case nilai >= 50:
		return "D"
	default:
		return "E"
	}
}

func middlewareNilai(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, isOk := r.BasicAuth()

		if !isOk {
			w.Write([]byte("Isi username dan password terlebih dahulu"))
			return
		}

		if username == "daffa" && password == "daffa123" {
			next.ServeHTTP(w, r)
			return
		}

		w.Write([]byte("Username atau password tidak sesuai"))
		return
	})
}

// Soal 2
func getNilaiMahasiswa(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", contentTypeJson)

	if r.Method == "GET" {
		dataNilai, err := json.Marshal(nilaiNilaiMahasiswa)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.WriteHeader(http.StatusOK)
		w.Write(dataNilai)
		return
	}

	http.Error(w, "Unknown error", http.StatusNotFound)
}

var nilaiNilaiMahasiswa []NilaiMahasiswa

const contentTypeJson string = "application/json"

func PostDataMahasiswa(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", contentTypeJson)
	var nilai NilaiMahasiswa

	if r.Method == "POST" {
		if r.Header.Get("Content-type") == contentTypeJson {
			decodedJson := json.NewDecoder(r.Body)
			if err := decodedJson.Decode(&nilai); err != nil {
				log.Fatal(err)
			}
			nilai.ID = uint(len(nilaiNilaiMahasiswa) + 1)
			nilai.IndeksNilai = getIndexNilai(int(nilai.Nilai))
		} else {
			Nama := r.PostFormValue("nama")
			Matkul := r.PostFormValue("matkul")
			Nilai, _ := strconv.Atoi(r.PostFormValue("nilai"))
			indexNilai := getIndexNilai(Nilai)
			Id := len(nilaiNilaiMahasiswa) + 1

			nilai = NilaiMahasiswa{
				ID:          uint(Id),
				Nilai:       uint(Nilai),
				IndeksNilai: indexNilai,
				Nama:        Nama,
				MataKuliah:  Matkul,
			}
		}
		nilaiNilaiMahasiswa = append(nilaiNilaiMahasiswa, nilai)

		dataMahasiswa, _ := json.Marshal(nilai)
		w.WriteHeader(http.StatusOK)
		w.Write(dataMahasiswa)
		return
	}
}

func main() {
	http.Handle("/tambahnilai", middlewareNilai(http.HandlerFunc(PostDataMahasiswa)))
	http.HandleFunc("/nilai", getNilaiMahasiswa)

	log.Println("Berjalan di port :10000")
	if err := http.ListenAndServe(":10000", nil); err != nil {
		log.Fatal(err.Error())
	}
}
