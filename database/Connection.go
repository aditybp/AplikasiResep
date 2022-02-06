package database

import (
	"aplikasiresep/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const DB_username = "root"
const DB_password = ""
const DB_name = "resep"
const DB_host = "127.0.0.1"
const DB_port = "3306"

var DB *gorm.DB

func Connection() {
	dsn := DB_username + ":" + DB_password + "@tcp" + "(" + DB_host + ":" + DB_port + ")/" + DB_name + "?" + "parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("error koneksi bermasalah !")
	}

	DB = db
	db.AutoMigrate(&models.BahanMakanan{})
	db.AutoMigrate(&models.KategoriMakanan{})
	db.AutoMigrate(&models.ResepMakanan{})
}
