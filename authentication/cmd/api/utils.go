package main

import (
	"context"
	"database/sql"
	"fmt"
)

func CreateUsersTableIfNotExists(db *sql.DB, ctx context.Context) error {
	stmt := `CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        email VARCHAR(255) NOT NULL,
        first_name VARCHAR(255),
        last_name VARCHAR(255),
        password VARCHAR(255) NOT NULL,
        user_active INT NOT NULL,
        created_at TIMESTAMP WITH TIME ZONE NOT NULL,
        updated_at TIMESTAMP WITH TIME ZONE NOT NULL
    )`

	_, err := db.ExecContext(ctx, stmt)
	return err
}

func DropUsersTableIfExist(db *sql.DB) error {
	ctx := context.Background()

	_, err := db.ExecContext(ctx, fmt.Sprintf("DROP TABLE IF EXISTS users"))
	return err
}
