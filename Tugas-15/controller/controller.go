package controller

import (
	"context"
	"fmt"
	"log"
	"time"
	"tugas-15/model"
	"tugas-15/obj"
)

const (
	table          string = "nilaimahasiswa"
	layoutDateTime string = "2006-01-02 15:04:05"
)

func GetAllNilai(context context.Context) ([]obj.NilaiMahasiswa, error) {
	var nilaiMahasiswa []obj.NilaiMahasiswa

	db, err := model.ConnectToMySQL()

	if err != nil {
		log.Fatal("Error connecting to database", err)
	}

	query := fmt.Sprintf("SELECT * FROM %v ORDER BY created_at DESC", table)

	rowQuery, err := db.QueryContext(context, query)

	if err != nil {
		log.Fatal("Error trying to query the database:", err)
	}

	for rowQuery.Next() {
		var nilaiSatuMahasiswa obj.NilaiMahasiswa
		var createdAt, updatedAt string

		if err := rowQuery.Scan(
			&nilaiSatuMahasiswa.ID,
			&nilaiSatuMahasiswa.Nama,
			&nilaiSatuMahasiswa.MataKuliah,
			&nilaiSatuMahasiswa.Nilai,
			&nilaiSatuMahasiswa.IndeksNilai,
			&createdAt,
			&updatedAt,
		); err != nil {
			return nil, err
		}

		// Change the time format
		nilaiSatuMahasiswa.CreatedAt, err = time.Parse(layoutDateTime, createdAt)

		if err != nil {
			log.Fatal(err)
		}

		nilaiSatuMahasiswa.UpdatedAt, err = time.Parse(layoutDateTime, updatedAt)

		if err != nil {
			log.Fatal(err)
		}

		nilaiMahasiswa = append(nilaiMahasiswa, nilaiSatuMahasiswa)
	}
	return nilaiMahasiswa, nil
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
	case nilai <= 49:
		return "E"
	default:
		return "EROR"
	}
}

func InsertNilai(context context.Context, nilai obj.NilaiMahasiswa) error {
	db, err := model.ConnectToMySQL()
	if err != nil {
		log.Fatal("Error connecting to database", err)
	}

	query := fmt.Sprintf("INSERT INTO %v (nama, matakuliah, nilai, indeksnilai, created_at, updated_at) VALUES('%v', '%v', %v, '%v', NOW(), NOW())", table, nilai.Nama, nilai.MataKuliah, nilai.Nilai, getIndexNilai(nilai.Nilai))

	_, err = db.QueryContext(context, query)

	if err != nil {
		log.Fatal("Error trying to query the database:", err)
	}
	return nil
}
