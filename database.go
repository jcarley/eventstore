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

func DbTime() (createdAt time.Time, updatedAt time.Time) {
	newTime := time.Now().UTC()
	createdAt = newTime
	updatedAt = newTime
	return
}
