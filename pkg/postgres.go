package pkg

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func ConnectPostgres() (*sql.DB, error) {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dbname := os.Getenv("DB_NAME")
	ssl := os.Getenv("DB_SSL")

	connStr := fmt.Sprintf("postgresql://%s:%s@%s/%s?sslmode=%s", user, password, host, dbname, ssl)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Panicln("error connStr", err.Error())
	}

	err = db.Ping()
	if err != nil {
		log.Panicln("error ping", err.Error())
	}

	err = migration(db)
	if err != nil {
		log.Panicln("error migration", err.Error())
	}

	return db, nil
}

func migration(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS companies (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL
	);
	`
	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("gagal membuat table companies")
	}

	query = `CREATE TABLE IF NOT EXISTS jobs (
    id UUID PRIMARY KEY,
    company_id UUID REFERENCES companies(id),
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`
	_, err = db.Exec(query)
	if err != nil {
		return fmt.Errorf("gagal membuat table jobs")
	}

	return err
}
