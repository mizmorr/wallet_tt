package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
)

func toPostgres() error {
	connStr := "postgres://postgres:post@db:5432/pgbounce"

	pool, err := pgxpool.Connect(context.Background(), connStr)
	if err != nil {
		return err
	}
	defer pool.Close()

	var greeting string
	err = pool.QueryRow(context.Background(), "SELECT 'Hello, world!'").Scan(&greeting)
	if err != nil {
		return err
	}

	fmt.Println(greeting)
	return nil
}

func viaBouncer() error {
	connStr := os.Getenv("DATABASE_URL")

	pool, err := pgxpool.Connect(context.Background(), connStr)
	if err != nil {
		return err
	}
	defer pool.Close()

	var greeting string
	err = pool.QueryRow(context.Background(), "SELECT 'Hello, world!'").Scan(&greeting)
	if err != nil {
		return err
	}

	err = pool.Ping(context.Background())
	if err != nil {
		return err
	}
	fmt.Println(greeting)
	return nil
}

func main() {
	http.HandleFunc("/postgr", func(w http.ResponseWriter, r *http.Request) {
		er := toPostgres()
		if er != nil {
			w.Write([]byte("database problem: " + er.Error()))
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.Write([]byte("everything is gOOD!"))
			w.WriteHeader(http.StatusOK)
		}
	})
	http.HandleFunc("/bounce", func(w http.ResponseWriter, r *http.Request) {
		er := viaBouncer()
		if er != nil {
			w.Write([]byte("server problem: " + er.Error()))
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			w.Write([]byte("all is good"))
			w.WriteHeader(http.StatusOK)
		}
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("listen error: %v", err)
	}
}
