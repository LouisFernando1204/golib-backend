package connection

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/LouisFernando1204/golang-restapi.git/internal/config"
	_ "github.com/lib/pq"
)

func GetDatabase(conf config.Database) *sql.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable Timezone=%s",
		conf.Host,
		conf.Port,
		conf.User,
		conf.Name,
		conf.Tz,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to open connection: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping connection: %v", err)
	}

	return db
}
