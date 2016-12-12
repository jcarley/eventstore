package eventstore

import (
	"time"

	"github.com/twinj/uuid"
)

type EventSource struct {
	ID         string
	SourceType string
	Version    int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func NewEventSource(sourceType string, version int) *EventSource {
	return &EventSource{
		ID:         uuid.NewV4().String(),
		SourceType: sourceType,
		Version:    version,
	}
}
