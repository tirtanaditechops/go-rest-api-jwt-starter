package config

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/denisenkom/go-mssqldb"
)

func ConnectDB() (*sql.DB, error) {
	dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s&trustservercertificate=true",
		os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))

	return sql.Open("sqlserver", dsn)
}
