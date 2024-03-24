package models

type Pertanyaan struct {
	ID            int64  `gorm:"primaryKey;autoIncrement" json:"id"`
	Pertanyaan    string `gorm:"not null" json:"pertanyaan"`
	Opsi_jawaban  string `gorm:"not null" json:"opsi_jawaban"`
	Jawaban_benar int64  `gorm:"not null" json:"jawaban_benar"`
	Id_quiz       int64  `gorm:"not null" json:"id_quiz"`
}
