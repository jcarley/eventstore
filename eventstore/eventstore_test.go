package eventstore

import (
	"testing"

	. "github.com/onsi/gomega"
)

func TestAddEventSource(t *testing.T) {
	RegisterTestingT(t)

	// streamName := fmt.Sprintf("%s-%s", "PayAsYouGoAccount", uuid.NewV4().String())
	// fmt.Println(streamName)
	// events := []interface{}{
	// examples.CreditAdded{5},
	// examples.PhoneCallCharged{20},
	// }

	// eventStore := NewPostgresEventStore()
	// eventStore.CreateNewStream(streamName, events)

	// eventSource := NewEventSource("SomeEventHappened", 1)
	// AddEventSource(eventSource)

	// Send register request
	// req, err := http.NewRequest("POST", "/api/v1/file/sync", strings.NewReader(body))
	// Expect(err).ShouldNot(HaveOccurred(), "Should be able to create a request")

	// req.Header.Add("Content-Type", "application/json")

	// w := httptest.NewRecorder()

	// controller := NewFileSyncController(renderer, database, pool)
	// controller.create(w, req)

	// Expect(w.Code).To(Equal(http.StatusCreated), "Should receive 201 status")
}
