package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"microservices/challenge-3/repository/postgres/sqlc"

	_ "github.com/lib/pq"
)

// DB cradential
var (
	// host     = os.Getenv("DB_HOST")
	// port     = os.Getenv("DB_PORT")
	// username = os.Getenv("DB_USERNAME")
	// password = os.Getenv("DB_PASSWORD")
	// dbname   = os.Getenv("DB_POSGRES")

	host     = "127.0.0.1"
	port     = 5432
	username = "root"
	password = "root"
	dbname   = "hacktiv"
)

func New() *sqlc.Queries {
	dbURL := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Jakarta", host, port, username, password, dbname)

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Cannot connect to DB:", err)
	}

	queries := sqlc.New(db)

	return queries
}
