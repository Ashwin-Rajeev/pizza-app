package cmd

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattes/migrate/source/file"
	"github.com/spf13/viper"

	_ "github.com/lib/pq"
)

const (
	// api env vars
	PortEvar    = "API_PORT"
	portDefault = ":3000"

	// postgres env vars
	pgEvarDb       = "POSTGRES_DB"
	pgEvarUser     = "POSTGRES_USER"
	pgEvarPassword = "POSTGRES_PASSWORD"
	pgEvarHost     = "POSTGRES_HOST"
	pgEvarPort     = "POSTGRES_PORT"
)

func mustReadPort() string {
	val := viper.GetString(PortEvar)
	if len(val) == 0 {
		log.Fatalf(
			"invalid port: %s",
			val,
		)
	}
	return val
}

// Prepare and connect PostgrSQL DB
func mustPrepareDB() *sql.DB {
	pgUser := viper.GetString(pgEvarUser)
	if len(pgUser) == 0 {
		log.Fatalf("invalid POSTGRES_USER")
	}
	pgPassword := viper.GetString(pgEvarPassword)
	if len(pgPassword) == 0 {
		log.Fatalf("invalid POSTGRES_PASSWORD")
	}
	dbHost := viper.GetString(pgEvarHost)
	if len(dbHost) == 0 {
		log.Fatalf("invalid POSTGRES_HOST")
	}
	dbPort := viper.GetString(pgEvarPort)
	if len(dbPort) == 0 {
		log.Fatalf("invalid POSTGRES_PORT")
	}
	dbName := viper.GetString(pgEvarDb)
	if len(dbPort) == 0 {
		log.Fatalf("invalid POSTGRES_DB")
	}
	connectionURL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable&search_path=public",
		pgUser, pgPassword, dbHost, dbPort, dbName)
	var db *sql.DB
	var err error
	for i := 0; i < 5; i++ {
		db, err = sql.Open("postgres", connectionURL)
		if err != nil {
			log.Println(err)
		}
		time.Sleep(time.Second)
	}
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file:///go/src/api/internal/data/migration",
		"postgres", driver)
	if err != nil {
		log.Fatal(err)
	}
	if err = m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatal(err)
	}
	log.Println("database migration completed")
	return db
}
