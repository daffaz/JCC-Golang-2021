package tugas8

import (
	"fmt"
	"strings"
)

// Deklarasi Pi

var PI = 3.14159265359

// Struct
// Soal 1
type SegitigaSamaSisi struct {
	Alas, Tinggi int
}

type PersegiPanjang struct {
	Panjang, Lebar int
}

type Tabung struct {
	JariJari, Tinggi float64
}

type Balok struct {
	Panjang, Lebar, Tinggi int
}

// Soal 2

type Phone struct {
	Name, Brand string
	Year        int
	Colors      []string
}

// Interface
type HitungBangunDatar interface {
	Luas() int
	Keliling() int
}

type HitungBangunRuang interface {
	Volume() float64
	LuasPermukaan() float64
}

// Soal 2
type TampilDataPhone interface {
	TampilData() string
}

// Function

// Segitiga
func (s SegitigaSamaSisi) Luas() int {
	return s.Alas * s.Tinggi / 2
}

func (s SegitigaSamaSisi) Keliling() int {
	return s.Alas * 3
}

// // Persegi Panjang
func (p PersegiPanjang) Luas() int {
	return p.Panjang * p.Lebar
}

func (p PersegiPanjang) Keliling() int {
	return (p.Panjang + p.Lebar) * 2
}

// Tabung
func (t Tabung) Volume() float64 {
	return PI * t.JariJari * t.JariJari * t.Tinggi
}

func (t Tabung) LuasPermukaan() float64 {
	return 2 * PI * t.JariJari * (t.JariJari + t.Tinggi)
}

// Balok
func (b Balok) Volume() float64 {
	return float64(b.Panjang * b.Lebar * b.Tinggi)
}

func (b Balok) LuasPermukaan() float64 {
	return float64(2*(b.Panjang*b.Lebar) + 2*(b.Panjang*b.Tinggi) + 2*(b.Lebar*b.Tinggi))
}

// Soal 2

func (p Phone) TampilData() string {
	result := fmt.Sprintf("name : %s\nbrand : %s\nyear : %d\ncolors : %s", p.Name, p.Brand, p.Year, strings.Join(p.Colors, ", "))
	return result
}

// Soal 3

func LuasPersegi(sisi int, status bool) interface{} {
	switch {
	case sisi > 1:
		if status {
			return sisi * sisi
		} else {
			return sisi
		}
	default:
		if status {
			return "Maaf anda belum menginput sisi dari persegi"
		} else {
			return nil
		}
	}
}
