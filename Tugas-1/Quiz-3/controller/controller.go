package controller

import (
	"context"
	"fmt"
	"log"
	"quiz3/config"
	"quiz3/model"
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
