package eventstore

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	db *sqlx.DB
)

func init() {
	ipAddress := "localhost"
	connectionString := fmt.Sprintf("postgres://admin:password@%s/eventstore_dev?sslmode=disable", ipAddress)
	// connection, err := sql.Open("postgres", connectionString)
	connection, err := sqlx.Connect("postgres", connectionString)
	if err != nil {
		panic(err)
	}
	db = connection
}

func GetDB() *sqlx.DB {
	return db
}

func DbTime() (createdAt time.Time, updatedAt time.Time) {
	newTime := time.Now().UTC()
	createdAt = newTime
	updatedAt = newTime
	return
}
