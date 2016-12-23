package providers

import (
	"errors"
	"os"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/jcarley/eventstore"
	"github.com/jmoiron/sqlx"
	"github.com/twinj/uuid"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

var (
	NoRecordsUpdatedError = errors.New("No records were updated.")
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	// log.SetFormatter(&log.JSONFormatter{})
	log.SetFormatter(&prefixed.TextFormatter{TimestampFormat: time.RFC3339})

	// Output to stderr instead of stdout, could also be a file.
	log.SetOutput(os.Stderr)

	// Only log the warning severity or above.
	log.SetLevel(log.DebugLevel)
}

type PostgresEventStore struct {
}

func NewPostgresEventStore() *PostgresEventStore {
	// log.SetLevel(log.PanicLevel)
	return &PostgresEventStore{}
}

func (this *PostgresEventStore) CreateNewStream(streamName string, events []eventstore.DomainEvent) error {
	// the uuid past in the first parameter is the primary key for the EventSource record.  The ID
	// for the Object is buried inside the streamName
	eventSource := eventstore.NewEventSourceWithID(uuid.NewV4().String(), streamName)

	createdAt, updatedAt := eventstore.TimeStamps()
	eventSource.CreatedAt = createdAt
	eventSource.UpdatedAt = updatedAt

	tx := this.startTransaction()
	this.addEventSource(eventSource, tx)
	this.commitTransaction(tx)

	// Expected version at the moment is going to be zero.
	return this.AppendEventsToStream(streamName, events, 0)
}

func (this *PostgresEventStore) AppendEventsToStream(streamName string, events []eventstore.DomainEvent, expectedVersion int) error {

	eventSource, err := this.getEventSourceByStreamName(streamName)
	if err != nil {
		return err
	}

	tx := this.startTransaction()
	for _, event := range events {
		this.saveEvent(eventSource.RegisterEvent(event), tx)
	}

	// update event source because the version will have changed
	err = this.updateEventSource(eventSource, tx)
	if err != nil {
		this.rollbackTransaction(tx)
		return err
	}
	this.commitTransaction(tx)

	return nil
}

func (this *PostgresEventStore) GetStream(streamName string, fromVersion int, toVersion int) ([]eventstore.DomainEvent, error) {

	db := eventstore.GetDB()
	events, err := this.getEventStream(streamName, fromVersion, toVersion, db)
	if err != nil {
		return nil, err
	}

	size := len(events)
	domainEvents := make([]eventstore.DomainEvent, size, size)
	for _, event := range events {
		domainEvent, err := event.Unwrap()
		if err != nil {
		}
		domainEvents = append(domainEvents, domainEvent)
	}

	return domainEvents, nil
}

func (this *PostgresEventStore) AddSnapshot(streamName string, snapShot interface{}) {
}

func (this *PostgresEventStore) GetLatestSnapshot(streamName string) (interface{}, error) {
	return nil, nil
}

func (this *PostgresEventStore) getEventSourceByStreamName(streamName string) (eventSource *eventstore.EventSource, err error) {

	db := eventstore.GetDB()
	statement := `select id, stream_name, version, created_at, updated_at
								from event_sources where stream_name = $1`

	eventSource = &eventstore.EventSource{}
	if err = db.Get(eventSource, statement, streamName); err != nil {
		eventSource = nil
	}

	return
}

func (this *PostgresEventStore) addEventSource(eventSource *eventstore.EventSource, tx *sqlx.Tx) error {

	statement := `insert into event_sources (id, stream_name, version, created_at, updated_at)
								values ($1, $2, $3, $4, $5)`

	log.WithFields(log.Fields{
		"ID":         eventSource.ID,
		"Version":    eventSource.Version,
		"StreamName": eventSource.StreamName,
	}).Info("Adding Event Source")

	tx.MustExec(statement,
		eventSource.ID,
		eventSource.StreamName,
		eventSource.Version,
		eventSource.CreatedAt.Format(time.RFC3339),
		eventSource.UpdatedAt.Format(time.RFC3339),
	)

	return nil
}

func (this *PostgresEventStore) updateEventSource(eventSource *eventstore.EventSource, tx *sqlx.Tx) error {
	statement := `update event_sources
								set version = $1,
									  updated_at = $2
								where id = $3`

	updatedAt := eventstore.NewFormattedDbTime()

	log.WithFields(log.Fields{
		"ID":         eventSource.ID,
		"Version":    eventSource.Version,
		"Updated At": updatedAt,
	}).Info("Updating Event Source")

	result := tx.MustExec(statement,
		eventSource.Version,
		updatedAt,
		eventSource.ID,
	)

	if rowsAffected, err := result.RowsAffected(); err != nil {
		return err
	} else if rowsAffected == 0 {

		log.WithFields(log.Fields{
			"ID":   eventSource.ID,
			"Type": "EventSource",
		}).Error(NoRecordsUpdatedError.Error())

		return NoRecordsUpdatedError
	}

	return nil
}

func (this *PostgresEventStore) getEventStream(streamName string, fromVersion int, toVersion int, db *sqlx.DB) ([]eventstore.EventWrapper, error) {

	statement := `SELECT ID, NAME, stream_name, VERSION, "sequence", time_stamp, "data"
								FROM events
								WHERE events.stream_name = $1
									AND events."version" > $2
									AND events."version" <= $3
								ORDER BY events."sequence"`

	events := []eventstore.EventWrapper{}
	err := db.Select(&events, statement, streamName, fromVersion, toVersion)

	if err != nil {
		return nil, err
	}

	return events, nil
}

func (this *PostgresEventStore) saveEvent(event *eventstore.EventWrapper, tx *sqlx.Tx) error {

	statement := `insert into events (id, time_stamp, name, version, stream_name, sequence, data)
							  values ($1, $2, $3, $4, $5, $6, $7)`

	log.WithFields(log.Fields{
		"ID":       event.ID,
		"Name":     event.Name,
		"Number":   event.EventNumber,
		"Sequence": event.Sequence,
	}).Info("Saving event")

	tx.MustExec(statement,
		event.ID,
		eventstore.NewFormattedDbTime(),
		event.Name,
		event.EventNumber,
		event.StreamName,
		event.Sequence,
		event.Event,
	)

	return nil
}

func (this *PostgresEventStore) startTransaction() *sqlx.Tx {
	db := eventstore.GetDB()
	return db.MustBegin()
}

func (this *PostgresEventStore) commitTransaction(tx *sqlx.Tx) error {
	return tx.Commit()
}

func (this *PostgresEventStore) rollbackTransaction(tx *sqlx.Tx) error {
	return tx.Rollback()
}
