package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"loketnadi-be-go/app"
	"loketnadi-be-go/pkg/database"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("❌ Gagal load .env:", err)
	}

	if err := database.ConnectDB(); err != nil {
		log.Fatal("❌ Gagal koneksi DB:", err)
	}

	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "3000"
	}

	log.Println("✅ Server jalan di port:", port)
	log.Fatal(app.Setup().Listen(":" + port))
}
