package eventstore

import (
	"fmt"
	"time"
)

type EventSource struct {
	ID         string    `json:"id"`
	SourceType string    `json:"sourceType"`
	Version    int       `json:"version"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

func NewEventSource(sourceType string) *EventSource {
	return &EventSource{
		SourceType: sourceType,
	}
}

func (this *EventSource) RegisterEvent(event interface{}) *EventWrapper {
	this.Version++
	return NewEventWrapper(event, this.Version, this.ID)
}

// **************************************************

type EventWrapper struct {
	ID            string      `json:"id"`
	Event         interface{} `json:"event"`
	EventSourceId string      `json:"eventSourceId"`
	EventNumber   int         `json:"eventNumber"`
}

func NewEventWrapper(event interface{}, eventNumber int, streamStateId string) *EventWrapper {
	return &EventWrapper{
		Event:         event,
		EventNumber:   eventNumber,
		EventSourceId: streamStateId,
		ID:            fmt.Sprintf("%s-%s", streamStateId, eventNumber),
	}
}
