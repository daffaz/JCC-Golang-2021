package controller

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/url"
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
		// var priceBook string

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

		// noRupiah := strings.Replace("Rp.100.000", "Rp.", "", -1)
		// noDot := strings.Replace(noRupiah, ".", "", -1)
		// noTrail := strings.Replace(noDot, ",-", "", -1)
		// buku.Price = noTrail

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

func formatCurrency(n int64) string {
	in := strconv.FormatInt(n, 10)
	numOfDigits := len(in)
	if n < 0 {
		numOfDigits--
	}
	numOfCommas := (numOfDigits - 1) / 3

	out := make([]byte, len(in)+numOfCommas)
	if n < 0 {
		in, out[0] = in[1:], '-'
	}

	for i, j, k := len(in)-1, len(out)-1, 0; ; i, j = i-1, j-1 {
		out[j] = in[i]
		if i == 0 {
			return string(out)
		}
		if k++; k == 3 {
			j, k = j-1, 0
			out[j] = '.'
		}
	}
}

func formatPrice(price int) string {
	fixedPrice := fmt.Sprintf("Rp.%v,-", formatCurrency(int64(price)))
	return fixedPrice
}

func validateYear(year int) (int, error) {
	if year >= 1980 && year <= 2021 {
		return year, nil
	}
	return -1, errors.New("tahun tidak boleh kurang dari 1980 dan tidak boleh lebih dari 2021")
}

func validateUrl(link string) (string, error) {
	_, err := url.ParseRequestURI(link)

	if err != nil {
		return "", errors.New("link tidak valid")
	}

	return link, nil
}

func InsertBuku(context context.Context, dataBuku model.Book) error {
	db, err := config.ConnectToMySQL()
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	tebalBuku, _ := strconv.Atoi(dataBuku.TotalPage)
	fixedPrice, _ := strconv.Atoi(dataBuku.Price)

	releaseYear, errYear := validateYear(dataBuku.ReleaseYear)
	fixedLink, errLink := validateUrl(dataBuku.ImageUrl)

	if errLink != nil && errYear != nil {
		log.Fatal("Error pada input URL gambar dan input tahun")
	}

	if errYear != nil {
		log.Fatal(errYear)
	}

	if errLink != nil {
		log.Fatal(errLink)
	}

	query := fmt.Sprintf("INSERT INTO %v (title, description, image_url, release_year, price, total_page, kategori_ketebalan, created_at, updated_at) VALUES ('%v', '%v', '%v', %v, '%v', '%v', '%v', NOW(), NOW())", tableName, dataBuku.Title, dataBuku.Description, fixedLink, releaseYear, formatPrice(fixedPrice), dataBuku.TotalPage, getKetebalanBuku(tebalBuku))

	_, err = db.QueryContext(context, query)

	if err != nil {
		log.Fatal("Error trying to insert to database:", err)
	}

	return nil
}

func UpdateBuku(context context.Context, dataBuku model.Book, idBuku string) error {
	db, err := config.ConnectToMySQL()

	if err != nil {
		log.Fatal("Error connecting to database", err)
	}
	tebalBuku, _ := strconv.Atoi(dataBuku.TotalPage)
	fixedPrice, _ := strconv.Atoi(dataBuku.Price)

	releaseYear, errYear := validateYear(dataBuku.ReleaseYear)
	fixedLink, errLink := validateUrl(dataBuku.ImageUrl)

	if errLink != nil && errYear != nil {
		log.Fatal("Error pada input URL gambar dan input tahun")
	}

	if errYear != nil {
		log.Fatal(errYear)
	}

	if errLink != nil {
		log.Fatal(errLink)
	}

	// title, description, image_url, release_year, price, total_page, kategori_ketebalan, created_at, updated_at
	query := fmt.Sprintf("UPDATE %v SET title='%v', description='%v', image_url='%v', release_year=%v, price = '%v', total_page = '%v', kategori_ketebalan = '%v', updated_at=NOW() WHERE id='%v'", tableName, dataBuku.Title, dataBuku.Description, fixedLink, releaseYear, formatPrice(fixedPrice), dataBuku.TotalPage, getKetebalanBuku(tebalBuku), idBuku)

	_, err = db.QueryContext(context, query)

	if err != nil {
		log.Fatal("Error trying to query the database:", err)
	}

	return nil
}

func DeleteBuku(context context.Context, idBuku string) error {
	db, err := config.ConnectToMySQL()

	if err != nil {
		log.Fatal("Error connecting to database", err)
	}

	query := fmt.Sprintf("DELETE FROM %v WHERE id = %v", tableName, idBuku)

	result, err := db.ExecContext(context, query)

	if err != nil && err != sql.ErrNoRows {
		return err
	}

	rowAffected, err := result.RowsAffected()

	log.Println(rowAffected)
	if rowAffected == 0 {
		return errors.New("can't delete, Id not found")
	}
	if err != nil {
		log.Fatal(err.Error())
	}

	return nil
}
