package controller

import (
	"context"
	"fmt"
	"log"
	"quiz3/config"
	"quiz3/model"
	"strconv"
	"time"
)

const (
	tableName  string = "books"
	dateLayout string = "2006-01-02 15:04:05"
)

func GetAllBooks(context context.Context) ([]model.Book, error) {
	var arrayBuku []model.Book

	db, err := config.ConnectToMySQL()
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	query := fmt.Sprintf("SELECT * FROM %v ORDER BY id DESC", tableName)

	rowQuery, err := db.QueryContext(context, query)
	if err != nil {
		log.Fatal("Error querying to database:", err)
	}

	for rowQuery.Next() {
		var buku model.Book
		var createdAt, updatedAt string

		if err := rowQuery.Scan(
			&buku.Id,
			&buku.Title,
			&buku.Description,
			&buku.ImageUrl,
			&buku.ReleaseYear,
			&buku.Price,
			&buku.TotalPage,
			&buku.KategoriKetebalan,
			&createdAt,
			&updatedAt,
		); err != nil {
			return nil, err
		}

		buku.CreatedAt, err = time.Parse(dateLayout, createdAt)
		if err != nil {
			log.Fatal("Error parsing created at:", err)
		}

		buku.UpdatedAt, err = time.Parse(dateLayout, updatedAt)
		if err != nil {
			log.Fatal("Error parsing updated at:", err)
		}

		arrayBuku = append(arrayBuku, buku)
	}
	return arrayBuku, nil
}

func getKetebalanBuku(halaman int) string {
	switch {
	case halaman >= 201:
		return "Tebal"
	case halaman >= 101 && halaman <= 200:
		return "Sedang"
	case halaman <= 100:
		return "Tipis"
	default:
		return "Undefined"
	}
}

func formatPrice(price int) string {
	fixedPrice := fmt.Sprintf("Rp.")
}

func InsertBuku(context context.Context, dataBuku model.Book) error {
	db, err := config.ConnectToMySQL()
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	tebalBuku, _ := strconv.Atoi(dataBuku.TotalPage)

	query := fmt.Sprintf("INSERT INTO %v (title, description, image_url, release_year, price, total_page, created_at, updated_at) VALUES ('%v', '%v', '%v', %v, %v, '%v', '%v', NOW(), NOW())", tableName, dataBuku.Title, dataBuku.Description, dataBuku.ImageUrl, dataBuku.ReleaseYear, dataBuku.Price, dataBuku.TotalPage, getKetebalanBuku(tebalBuku))

}
