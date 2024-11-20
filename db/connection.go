package db

import (
	"database/sql"
	"fmt"
	"log"
	"viewlens/config"

	_ "github.com/denisenkom/go-mssqldb"
)

func Connect(cfg config.Config) *sql.DB {
	connString := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s",
		cfg.DbUser, cfg.DbPassword, cfg.DbServer, cfg.DbPort, cfg.DbDatabase)

	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	return db
}
