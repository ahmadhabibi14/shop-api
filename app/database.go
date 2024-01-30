package app

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/ahmadhabibi14/shop-api/helper"
	_ "github.com/go-sql-driver/mysql"
)

func NewDB() *sql.DB {
	DbDriver := "mysql"
	DbHost := os.Getenv("MARIADB_HOST")
	DbPort := os.Getenv("MARIADB_PORT")
	DbName := os.Getenv("MARIADB_DATABASE")
	DbUser := os.Getenv("MARIADB_USER")
	DbPassword := os.Getenv("MARIADB_PASSWORD")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", DbUser, DbPassword, DbHost, DbPort, DbName)
	db, err := sql.Open(DbDriver, dsn)

	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxIdleTime(10 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)
	return db
}
