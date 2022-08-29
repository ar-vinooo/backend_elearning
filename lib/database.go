package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DB() (db *gorm.DB) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/e_learning?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()

	if err != nil {
		panic(err)
	}

	sqlDB.SetMaxIdleConns(10)

	sqlDB.SetMaxOpenConns(100)

	return db
}
