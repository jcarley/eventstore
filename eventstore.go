package eventstore

import (
	"time"

	"github.com/twinj/uuid"
)

type EventStore interface {
	CreateNewStream(streamName string, changes []interface{})
	AppendEventsToStream(streamName string, changes []interface{}, expectedVersion int) error
	GetStream(streamName string, fromVersion int, toVersion int) ([]interface{}, error)
	AddSnapshot(streamName string, snapShot interface{})
	GetLatestSnapshot(streamName string) (interface{}, error)
}

type PostgresEventStore struct {
}

func NewPostgresEventStore() *PostgresEventStore {
	return &PostgresEventStore{}
}

func (this *PostgresEventStore) CreateNewStream(streamName string, events []interface{}) {
	eventSource := NewEventSource(streamName)
	eventSource.ID = uuid.NewV4().String()
	eventSource.Version = 0

	createdAt, updatedAt := DbTime()
	eventSource.CreatedAt = createdAt
	eventSource.UpdatedAt = updatedAt

	this.addEventSource(eventSource)

	this.AppendEventsToStream(streamName, events, 0)
}

func (this *PostgresEventStore) AppendEventsToStream(streamName string, events []interface{}, expectedVersion int) error {

	eventSource := this.getEventSourceByStreamName(streamName)

	for _, event := range events {
		this.saveEvent(event, eventSource.ID)
	}
}

func (this *PostgresEventStore) addEventSource(eventSource *EventSource) {

	db := GetDB()
	statement := `insert into event_sources (id, source_type, version, created_at, updated_at)
								values ($1, $2, $3, $4, $5)`

	stmt, err := db.Prepare(statement)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	stmt.QueryRow(
		eventSource.ID,
		eventSource.SourceType,
		eventSource.Version,
		eventSource.CreatedAt.Format(time.RFC3339),
		eventSource.UpdatedAt.Format(time.RFC3339),
	)

	return
}
