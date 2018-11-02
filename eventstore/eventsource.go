package eventstore

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/jcarley/eventstore/helper/jsonutil"
	"github.com/mitchellh/mapstructure"
	"github.com/twinj/uuid"
)

type EventSource struct {
	ID         string    `json:"id" db:"id"`
	StreamName string    `json:"streamName" db:"stream_name"`
	Version    int       `json:"version" db:"version"`
	CreatedAt  time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt  time.Time `json:"updatedAt" db:"updated_at"`
}

func NewEventSource(streamName string) *EventSource {
	return &EventSource{
		StreamName: streamName,
		Version:    0,
	}
}

func NewEventSourceWithID(id string, streamName string) *EventSource {
	return &EventSource{
		ID:         id,
		StreamName: streamName,
		Version:    0,
	}
}

func (this *EventSource) RegisterEvent(event DomainEvent) *EventWrapper {
	this.Version++
	return NewEventWrapper(event, this.Version, this.StreamName, this.Version)
}

type EventWrapper struct {
	ID          string    `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Event       string    `json:"event" db:"data"`
	StreamName  string    `json:"streamName" db:"stream_name"`
	EventNumber int       `json:"eventNumber" db:"version"`
	Sequence    int       `json:"sequence" db:"sequence"`
	Timestamp   time.Time `json:"time_stamp" db:"time_stamp"`
}

func NewEventWrapper(event DomainEvent, eventNumber int, streamName string, sequence int) *EventWrapper {

	data, _ := jsonutil.EncodeJSONToString(event)

	return &EventWrapper{
		ID:          uuid.NewV4().String(),
		Name:        fmt.Sprintf("%T", event),
		Event:       data,
		StreamName:  streamName,
		EventNumber: eventNumber,
		Sequence:    sequence,
	}
}

func (this *EventWrapper) Unwrap() (DomainEvent, error) {

	reader := strings.NewReader(this.Event)
	var output map[string]interface{}

	if err := jsonutil.DecodeJSONFromReader(reader, &output); err != nil {
		return (DomainEvent)(nil), err
	}

	eventType := MakeInstance(this.Name)

	if err := mapstructure.Decode(output, eventType); err != nil {
		return (DomainEvent)(nil), err
	}

	event, ok := eventType.(DomainEvent)

	if ok {
		return event, nil
	} else {
		return (DomainEvent)(nil), errors.New("")
	}
}
