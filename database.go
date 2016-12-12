package eventstore

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

var (
	db *sql.DB
)

func init() {
	ipAddress := "localhost"
	connectionString := fmt.Sprintf("postgres://admin:password@%s/eventstore_dev?sslmode=disable", ipAddress)
	connection, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}
	db = connection
}

func GetDB() *sql.DB {
	return db
}

func DbTime() (createdAt string, updatedAt string) {
	newTime := time.Now().UTC().Format(time.RFC3339)
	createdAt = newTime
	updatedAt = newTime
	return
}
