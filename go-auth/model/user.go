package models

type User struct {
	ID       int64  `gorm:"primaryKey;autoIncrement" json:"id"`
	Nama     string `gorm:"not null" json:"nama"`
	Email    string `gorm:"not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
	Role     string `gorm:"not null" json:"role"`
}
