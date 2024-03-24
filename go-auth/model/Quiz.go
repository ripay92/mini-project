package models

import "time"

type Quiz struct {
	ID           int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Judul        string    `gorm:"not null" json:"judul"`
	Deskripsi    string    `gorm:"not null" json:"deskripsi"`
	WaktuMulai   time.Time `gorm:"not null" json:"waktu_mulai"`
	WaktuSelesai time.Time `gorm:"not null" json:"waktu_selesai"`
}
