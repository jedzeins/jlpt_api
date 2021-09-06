package repositories

import (
	"database/sql"
	"fmt"
)

func InitUsersTable(DB *sql.DB) error {
	_, err := DB.Exec(
		`
		CREATE TABLE IF NOT EXISTS users 
			(
				id SERIAL PRIMARY KEY,
				username VARCHAR(20) NOT NULL UNIQUE, 
				password VARCHAR(255) NOT NULL, 
				active BOOL NOT NULL DEFAULT true, 
				isAdmin BOOL NOT NULL DEFAULT false,
				created_date timestamp without time zone NOT NULL DEFAULT (current_timestamp AT TIME ZONE 'UTC'), 
				last_login timestamp without time zone DEFAULT (current_timestamp AT TIME ZONE 'UTC'), 
				last_modified_date timestamp without time zone 
			)
		`,
	)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return err

}

func InitSessionsTable(DB *sql.DB) error {
	_, err := DB.Exec(
		`
		CREATE TABLE IF NOT EXISTS sessions 
			(
				id SERIAL PRIMARY KEY,
				owner_id integer NOT NULL,
				created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
				FOREIGN KEY (owner_id) REFERENCES users (id) ON DELETE CASCADE
			)
		`,
	)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return err

}
