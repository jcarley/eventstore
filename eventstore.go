package eventstore

type EventStore interface {
	CreateNewStream(streamName string, events []DomainEvent) error
	AppendEventsToStream(streamName string, events []DomainEvent, expectedVersion int) error
	GetStream(streamName string, fromVersion int, toVersion int) ([]DomainEvent, error)
	AddSnapshot(streamName string, snapShot interface{})
	GetLatestSnapshot(streamName string) (interface{}, error)
}
