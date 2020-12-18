package config

import (
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func GetDB() (db *gorm.DB, err error) {
	//	db, err = sql.Open("mssql", "server=185.94.98.228;user id=rose;password=z6prT9@5;database=iranvar;")
	//db, err = gorm.Open("mssql", "server=185.94.98.228;user id=rose;password=z6prT9@5;database=iranvar;")
	dsn := "server=185.94.98.228;user id=rose;password=z6prT9@5;database=iranvar;"
	db, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	return
}
