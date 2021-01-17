package config

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/gorm"
)

func GetDB() (db *gorm.DB, err error) {
	//	db, err = sql.Open("mssql", "server=185.94.98.228;user id=rose;password=z6prT9@5;database=iranvar;")
	//db, err = gorm.Open("mssql", "server=185.94.98.228;user id=rose;password=z6prT9@5;database=iranvar;")
	// dsn := "server=185.94.98.228;user id=rose;password=z6prT9@5;database=iranvar;"
	// db, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{})

	dsn := "host=localhost user=postgres password=8elmira8 dbname=zapood port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// db, err := gorm.Open(postgres.New(postgres.Config{
	// 	DSN: "host=localhost user=postgres password=8elmira8 DB.name=zapood port=5432 sslmode=disable TimeZone=Asia/Shanghai", // data source name, refer https://github.com/jackc/pgx
	// 	PreferSimpleProtocol: true,                                                                                                             // disables implicit prepared statement usage. By default pgx automatically uses the extended protocol
	// }), &gorm.Config{})

	return
}
