package main

import (
	{{if eq .Database "PostgreSQL"}}
	_ "github.com/lib/pq"
	"database/sql"
	{{else if eq .Database "MongoDB"}}
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"context"
	"time"
	{{else if eq .Database "MySQL"}}
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	{{else if eq .Database "SQLite"}}
	_ "github.com/mattn/go-sqlite3"
	"database/sql"
	{{end}}
	"fmt"
)

func initDB() {
	{{if eq .Database "PostgreSQL"}}
	dsn := "host=localhost user=root password=secret dbname={{.ProjectName}} sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	{{else if eq .Database "MongoDB"}}
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	db := client.Database("{{.ProjectName}}")
	{{else if eq .Database "MySQL"}}
	dsn := "root:password@tcp(127.0.0.1:3306)/{{.ProjectName}}"
	db, err := sql.Open("mysql", dsn)
	{{else if eq .Database "SQLite"}}
	dsn := "{{.ProjectName}}.db"
	db, err := sql.Open("sqlite3", dsn)
	{{else}}
	fmt.Println("ℹ️ No database selected.")
	return
	{{end}}

	if err != nil {
		fmt.Println("❌ Database connection failed:", err)
		return
	}

	fmt.Println("✅ Database connected successfully!")
}
