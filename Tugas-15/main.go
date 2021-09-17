package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"tugas-15/controller"
	"tugas-15/obj"
	"tugas-15/utils"

	"github.com/julienschmidt/httprouter"
)

func GetNilaiMahasiswa(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	context, cancel := context.WithCancel(context.Background())

	defer cancel()
	nilaiMahasiswa, err := controller.GetAllNilai(context)

	if err != nil {
		utils.ResponseToJSON(w, err, http.StatusBadRequest)
		return
	}

	utils.ResponseToJSON(w, nilaiMahasiswa, http.StatusOK)
}

func PostNilaiMahasiswa(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Header.Get("Content-type") != "application/json" {
		utils.ResponseToJSON(w, "Only accepting application/json POST method", http.StatusBadRequest)
		return
	}

	context, cancel := context.WithCancel(context.Background())

	defer cancel()

	var nilai obj.NilaiMahasiswa
	if err := json.NewDecoder(r.Body).Decode(&nilai); err != nil {
		utils.ResponseToJSON(w, err, http.StatusBadRequest)
	}

	if err := controller.InsertNilai(context, nilai); err != nil {
		utils.ResponseToJSON(w, err, http.StatusInternalServerError)
		return
	}

	response := map[string]string{
		"message": "Post Success",
	}

	utils.ResponseToJSON(w, response, http.StatusCreated)
}

func UpdateNilaiMahasiswa(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	if r.Header.Get("Content-type") != "application/json" {
		utils.ResponseToJSON(w, "Only accepting application/json POST method", http.StatusBadRequest)
		return
	}

	context, cancel := context.WithCancel(context.Background())
	defer cancel()

	var nilai obj.NilaiMahasiswa
	if err := json.NewDecoder(r.Body).Decode(&nilai); err != nil {
		utils.ResponseToJSON(w, err, http.StatusBadRequest)
	}

	var idMahasiswa = param.ByName("id")

	if err := controller.UpdateNilai(context, nilai, idMahasiswa); err != nil {
		utils.ResponseToJSON(w, err, http.StatusInternalServerError)
		return
	}

	response := map[string]string{
		"message": "Update Succes",
	}

	utils.ResponseToJSON(w, response, http.StatusCreated)
}

func DeleteNilaiMahasiswa(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	context, cancel := context.WithCancel(context.Background())
	defer cancel()

	var idMahasiswa = param.ByName("id")
	if err := controller.DeleteNilaiMahasiswa(context, idMahasiswa); err != nil {
		utils.ResponseToJSON(w, err, http.StatusInternalServerError)
		return
	}

	res := map[string]string{
		"status": "Delete Success",
	}
	utils.ResponseToJSON(w, res, http.StatusOK)
}

func main() {
	router := httprouter.New()

	router.GET("/nilaimahasiswa", GetNilaiMahasiswa)
	router.POST("/nilaimahasiswa", PostNilaiMahasiswa)
	router.PUT("/nilaimahasiswa/:id", UpdateNilaiMahasiswa)
	router.DELETE("/nilaimahasiswa/:id", DeleteNilaiMahasiswa)

	log.Println("Running in port :10000")
	log.Fatal(http.ListenAndServe(":10000", router))
}
