package providers

import (
	"time"

	"github.com/jcarley/eventstore"
	"github.com/twinj/uuid"
)

// CreateNewStream(streamName string, events []DomainEvent)
// AppendEventsToStream(streamName string, events []DomainEvent, expectedVersion int) error
// GetStream(streamName string, fromVersion int, toVersion int) ([]DomainEvent, error)
// AddSnapshot(streamName string, snapShot interface{})
// GetLatestSnapshot(streamName string) (interface{}, error)

type PostgresEventStore struct {
}

func NewPostgresEventStore() *PostgresEventStore {
	return &PostgresEventStore{}
}

func (this *PostgresEventStore) CreateNewStream(streamName string, events []eventstore.DomainEvent) {
	eventSource := eventstore.NewEventSource(streamName)
	eventSource.ID = uuid.NewV4().String()
	eventSource.Version = 0

	createdAt, updatedAt := eventstore.DbTime()
	eventSource.CreatedAt = createdAt
	eventSource.UpdatedAt = updatedAt

	this.addEventSource(eventSource)

	this.AppendEventsToStream(streamName, events, 0)
}

func (this *PostgresEventStore) AppendEventsToStream(streamName string, events []eventstore.DomainEvent, expectedVersion int) error {

	eventSource, err := this.getEventSourceByStreamName(streamName)
	if err != nil {
		return err
	}

	for _, event := range events {
		this.saveEvent(eventSource.RegisterEvent(event))
	}

	return nil
}

func (this *PostgresEventStore) getEventSourceByStreamName(streamName string) (eventSource *eventstore.EventSource, err error) {

	db := eventstore.GetDB()
	statement := `select id source_type version created_at updated_at
								from event_sources where source_type = $1`
	eventSource = &eventstore.EventSource{}
	if err = db.Get(eventSource, statement, streamName); err != nil {
		eventSource = nil
	}

	return
}

func (this *PostgresEventStore) addEventSource(eventSource *eventstore.EventSource) {

	db := eventstore.GetDB()
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

func (this *PostgresEventStore) saveEvent(event eventstore.EventWrapper) {

}
