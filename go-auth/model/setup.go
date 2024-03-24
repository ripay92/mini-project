package models

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
func ConnectDatabase(){
	db, err:= gorm.Open(mysql.Open("root:@tcp(localhost:3306)/PNM"))
	if err != nil{
		panic(err)
	}
	
	DB = db
	log.Println("Berhasil terhubung ke db")
}