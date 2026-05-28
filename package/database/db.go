package database

import (
	"log"
	"main/package/config"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go.yaml.in/yaml/v3"
)

type NewPostgresDb struct {
	db *sqlx.DB
}

func (pg *NewPostgresDb) Connect() error {
	var cfg config.Config
	file, err := os.ReadFile("config.yml")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	yaml.Unmarshal(file, &cfg)
	db, err := sqlx.Connect("postgres", cfg.DBCfg.Url)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	pg.db = db
	return nil
}
