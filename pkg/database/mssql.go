package database

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
)

var DB *sql.DB

func ConnectDB() error {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf(
		"sqlserver://%s:%s@%s:%s?database=%s&encrypt=disable&trustservercertificate=true",
		user, pass, host, port, dbname,
	)

	var err error
	DB, err = sql.Open("sqlserver", dsn)
	if err != nil {
		return err
	}

	// Connection pool settings
	DB.SetMaxOpenConns(100)                // maks 50 koneksi total
	DB.SetMaxIdleConns(50)                 // simpan 25 idle untuk reuse
	DB.SetConnMaxLifetime(5 * time.Minute) // koneksi ganti tiap 5 menit

	return DB.Ping()
}
