package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

func GormConnect() *gorm.DB {
	DBMS := "sqlite3"
	DBNAME := "wanted.db"

	db, err := gorm.Open(DBMS, DBNAME)

	if err != nil {
		//DBが存在しない、または接続不能
		panic(err.Error())
	}

	db.LogMode(true)

	return db
}
