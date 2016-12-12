package eventstore

import (
	"fmt"
	"reflect"
)

// type DomainEvent struct {
// }

type EventSourcedAggregate struct {
	changes []interface{}
	version int
}

func NewEventSourcedAggregate() *EventSourcedAggregate {
	return &EventSourcedAggregate{
		changes: make([]interface{}, 0, 5),
		version: 0,
	}
}

func (this *EventSourcedAggregate) Version() int {
	return this.version
}

func (this *EventSourcedAggregate) Changes() []interface{} {
	return this.changes
}

func (this *EventSourcedAggregate) Apply(aggregate interface{}, change interface{}) {
	eventTypeName := reflect.TypeOf(change).Name()
	methodName := fmt.Sprintf("When%s", eventTypeName)
	this.Invoke(aggregate, methodName, change)

	this.version++
}

func (this *EventSourcedAggregate) Causes(aggregate interface{}, change interface{}) {
	this.changes = append(this.changes, change)
	this.Apply(aggregate, change)
}

func (this *EventSourcedAggregate) Invoke(any interface{}, name string, args ...interface{}) {
	inputs := make([]reflect.Value, len(args))
	for i, _ := range args {
		inputs[i] = reflect.ValueOf(args[i])
	}
	reflect.ValueOf(any).MethodByName(name).Call(inputs)
}
