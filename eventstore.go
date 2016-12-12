package eventstore

type EventStore interface {
	CreateNewStream()
	AppendEventsToStream()
	GetStream()
	AddSnapshot()
	GetLatestSnapshot()
}

func AddEventSource(eventSource *EventSource) {

	db := GetDB()
	statement := `insert into event_sources (id, source_type, version, created_at, updated_at)
								values ($1, $2, $3, $4, $5)`

	stmt, err := db.Prepare(statement)
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	createdAt, updatedAt := DbTime()

	stmt.QueryRow(
		eventSource.ID,
		eventSource.SourceType,
		eventSource.Version,
		createdAt,
		updatedAt,
	)

	return
}
