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

func NewDbTime() time.Time {
	return time.Now().UTC()
}

func NewFormattedDbTime() string {
	t := NewDbTime()
	return FormatDbTime(t)
}

func FormatDbTime(t time.Time) string {
	return t.Format(time.RFC3339)
}

func TimeStamps() (createdAt time.Time, updatedAt time.Time) {
	newTime := NewDbTime()
	createdAt = newTime
	updatedAt = newTime
	return
}
