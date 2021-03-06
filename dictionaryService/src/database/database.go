package database

import (
	"database/sql"
	"fmt"
	"time"
)

var (
	host         = "database" // if using docker compose, should be name of service, else "localhost"
	port         = 5432
	user         = "postgres"
	password     = "password"
	databaseName = "postgres"
	DB           *sql.DB
)

func Initialize() error {
	time.Sleep(3 * time.Second)
	dns := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, databaseName)

	var err error
	DB, err = sql.Open("postgres", dns)
	if err != nil {
		panic(err)
	}

	err = DB.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("DB CONNECTED")

	return err
}

func CheckDB() bool {
	var (
		err    error
		exists = true
		id     int8
	)

	// err = DB.QueryRow("SELECT EXISTS(SELECT * FROM information_schema.tables WHERE table_schema = $1 AND table_name = $2);", "public", "words").Scan(&exists)

	err = DB.QueryRow("SELECT id FROM words WHERE id = 1").Scan(&id)

	if err != nil || id != 1 {
		exists = false
	}

	return exists
}
