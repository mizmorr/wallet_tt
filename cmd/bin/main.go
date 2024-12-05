package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {
	connStr := "postgres://postgres:post@localhost:6432/pgbounce"

	pool, err := pgxpool.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatalf("unable to connect to database: %v\n", err)
	}
	defer pool.Close()

	var greeting string
	err = pool.QueryRow(context.Background(), "SELECT 'Hello, world!'").Scan(&greeting)
	if err != nil {
		log.Fatalf("query failed: %v\n", err)
	}

	fmt.Println(greeting)
}
