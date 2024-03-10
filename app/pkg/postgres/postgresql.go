package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

//type PostgresDB struct {
//	db *sql.DB
//}

func NewPostgresDB() (*sql.DB, error) {
	cfg := &Config{
		Host:     "localhost",
		Port:     "5440",
		Username: "postgres",
		Password: "secret",
		DBName:   "coffemapper",
		SSLMode:  "disable",
	}

	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBName, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
