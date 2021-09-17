package obj

import "time"

type NilaiMahasiswa struct {
	ID          uint      `json:"id"`
	Nama        string    `json:"nama"`
	MataKuliah  string    `json:"matakuliah"`
	Nilai       int       `json:"nilai"`
	IndeksNilai string    `json:"indeksnilai"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
