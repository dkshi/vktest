package repository

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Host     string
	DBName   string
	Port     string
	Username string
	Password string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sqlx.DB, error) {
	connStr := fmt.Sprintf("host=%s dbname=%s port=%s user=%s password=%s sslmode=%s", cfg.Host, cfg.DBName, cfg.Port, cfg.Username, cfg.Password, cfg.SSLMode)
	for {
		db, err := sqlx.Connect("postgres", connStr)
		if err == nil {
			InitSchema(db)
			if err == nil {
				return db, nil
			}
		}
		logrus.Printf("Failed to connect to the database: %v. Retrying...", err)
		time.Sleep(5 * time.Second)
	}
}

func InitSchema(db *sqlx.DB) {
	schema := `CREATE TABLE IF NOT EXISTS actors (
		actor_id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		gender VARCHAR(10) NOT NULL,
		birth_date DATE NOT NULL
	);
	
	CREATE TABLE IF NOT EXISTS films (
		film_id SERIAL PRIMARY KEY,
		title VARCHAR(150) NOT NULL,
		description TEXT,
		release_date DATE,
		rating NUMERIC(3,1) CHECK (rating >= 0 AND rating <= 10)
	);
	
	CREATE TABLE IF NOT EXISTS actor_film (
		actor_id INT REFERENCES actors(actor_id),
		film_id INT REFERENCES films(film_id),
		PRIMARY KEY (actor_id, film_id)
	);
	
	CREATE TABLE IF NOT EXISTS admins (
		admin_id SERIAL PRIMARY KEY,
		adminname VARCHAR(255) NOT NULL UNIQUE,
		password VARCHAR(255) NOT NULL
	);`
	db.MustExec(schema)
}