package eventstore

import "reflect"

type EventStore interface {
	CreateNewStream(streamName string, events []DomainEvent) error
	AppendEventsToStream(streamName string, events []DomainEvent, expectedVersion int) error
	GetStream(streamName string, fromVersion int, toVersion int) ([]DomainEvent, error)
	AddSnapshot(streamName string, snapShot interface{})
	GetLatestSnapshot(streamName string) (interface{}, error)
}

var typeRegistry = make(map[string]reflect.Type)

func RegisterType(typedNil interface{}) {
	t := reflect.TypeOf(typedNil).Elem()
	typeRegistry[t.String()] = t
}

func MakeInstance(name string) interface{} {
	t := typeRegistry[name]
	vp := reflect.New(t)
	return vp.Interface()
}
