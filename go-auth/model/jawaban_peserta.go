package models

type Jawaban_Peserta struct {
	ID              int64 `gorm:"primaryKey;autoIncrement" json:"id"`
	Id_user         int64 `gorm:"not null" json:"id_user"`
	Id_quiz         int64 `gorm:"not null" json:"id_quiz"`
	Id_pertanyaan   int64 `gorm:"not null" json:"id_pertanyaan"`
	Jawaban_Peserta int64 `gorm:"not null" json:"jawaban_peserta"`
	Skor            int64 `gorm:"not null" json:"skor"`
}
