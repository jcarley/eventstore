package examples

import "github.com/jcarley/eventstore"

type InMemoryEventStore struct {
}

func NewInMemoryEventStore() *InMemoryEventStore {
	return &InMemoryEventStore{}
}

func (this *InMemoryEventStore) CreateNewStream(streamName string, events []eventstore.DomainEvent) {
}

func (this *InMemoryEventStore) AppendEventsToStream(streamName string, events []eventstore.DomainEvent, expectedVersion int) error {
	return nil
}

func (this *InMemoryEventStore) GetStream(streamName string, fromVersion int, toVersion int) ([]eventstore.DomainEvent, error) {
	return nil, nil
}

func (this *InMemoryEventStore) AddSnapshot(streamName string, snapShot interface{}) {
}

func (this *InMemoryEventStore) GetLatestSnapshot(streamName string) (interface{}, error) {
	return nil, nil
}
