package eventstore

import (
	"fmt"
	"time"

	"github.com/twinj/uuid"
)

type EventSource struct {
	ID         string    `json:"id" db:"id"`
	SourceType string    `json:"sourceType" db:"source_type"`
	Version    int       `json:"version" db:"version"`
	CreatedAt  time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt  time.Time `json:"updatedAt" db:"updated_at"`
}

func NewEventSource(sourceType string) *EventSource {
	return &EventSource{
		SourceType: sourceType,
		Version:    0,
	}
}

func NewEventSourceWithID(id string, sourceType string) *EventSource {
	return &EventSource{
		ID:         id,
		SourceType: sourceType,
		Version:    0,
	}
}

func (this *EventSource) RegisterEvent(event DomainEvent) *EventWrapper {
	this.Version++
	return NewEventWrapper(event, this.Version, this.ID, this.Version)
}

// **************************************************

// "id" uuid NOT NULL,
// "time_stamp" timestamp(6) NOT NULL,
// "name" varchar NOT NULL COLLATE "default",
// "version" varchar NOT NULL COLLATE "default",
// "event_source_id" uuid NOT NULL,
// "sequence" int8,
// "data" json NOT NULL

type EventWrapper struct {
	ID            string      `json:"id" db:"id"`
	Name          string      `json:"name" db:"name"`
	Event         interface{} `json:"event" db:"data"`
	EventSourceId string      `json:"eventSourceId" db:"event_source_id"`
	EventNumber   int         `json:"eventNumber" db:"version"`
	Sequence      int         `json:"sequence" db:"sequence"`
}

func NewEventWrapper(event DomainEvent, eventNumber int, streamStateId string, sequence int) *EventWrapper {
	return &EventWrapper{
		ID:            uuid.NewV4().String(),
		Name:          fmt.Sprintf("%T", event),
		Event:         event,
		EventSourceId: streamStateId,
		EventNumber:   eventNumber,
		Sequence:      sequence,
	}
}
