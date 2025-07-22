package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	dsn := "host=localhost port=5432 user=louisfernando dbname=golang_restapi sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Gagal membuka koneksi: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("Gagal ping ke database: %v", err)
	}

	fmt.Println("Berhasil terhubung ke database!")

	usersToSeed := []map[string]string{
		{"email": "admin@example.com", "password": "password123"},
		{"email": "user1@example.com", "password": "password456"},
	}

	for _, userData := range usersToSeed {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userData["password"]), bcrypt.DefaultCost)
		if err != nil {
			log.Fatalf("Gagal hash password: %v", err)
		}

		query := `INSERT INTO users (id, email, password) VALUES ($1, $2, $3)`

		_, err = db.Exec(query, uuid.NewString(), userData["email"], string(hashedPassword))
		if err != nil {
			log.Printf("Gagal menyisipkan data untuk %s: %v", userData["email"], err)
		} else {
			fmt.Printf("Berhasil menyisipkan data untuk: %s\n", userData["email"])
		}
	}

	fmt.Println("Proses seeding data dummy selesai.")
}
