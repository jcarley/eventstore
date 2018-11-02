package eventstore

import (
	"fmt"
	"testing"

	"github.com/jcarley/eventstore/helper/jsonutil"
	. "github.com/onsi/gomega"
	"github.com/twinj/uuid"
)

type EmailSent struct {
	ToField   string `json:"to" mapstructure:"to"`
	FromField string `json:"from" mapstructure:"from"`
	Message   string `json:"message" mapstructure:"message"`
}

func (this EmailSent) Apply(es EventSourceAggregate) {
}

func init() {
	RegisterType((*EmailSent)(nil))
}

func TestWrappingAnEvent(t *testing.T) {
	RegisterTestingT(t)

	emailSent := EmailSent{
		ToField:   "jeff.carley@gmail.com",
		FromField: "john.doe@gmail.com",
		Message:   "This is the message",
	}

	eventWrapper := NewEventWrapper(emailSent, 1, "streamName", 5)

	expectedEventJson, _ := jsonutil.EncodeJSONToString(emailSent)
	fmt.Println(expectedEventJson)

	Expect(eventWrapper.Event).To(MatchJSON(expectedEventJson))
}

func TestUnwrappingAnEvent(t *testing.T) {
	RegisterTestingT(t)

	eventJson := `{"to":"jeff.carley@gmail.com","from":"john.doe@gmail.com","message":"This is the message"}`
	streamName := "streamName"
	eventNumber := 5
	sequence := 1

	eventWrapper := EventWrapper{
		ID:          uuid.NewV4().String(),
		Name:        fmt.Sprintf("%T", EmailSent{}),
		Event:       eventJson,
		StreamName:  streamName,
		EventNumber: eventNumber,
		Sequence:    sequence,
	}

	emailSentEvent, err := eventWrapper.Unwrap()

	Expect(err).ShouldNot(HaveOccurred())
	Expect(emailSentEvent).ToNot(BeNil())

	fmt.Printf("%#v", emailSentEvent)
}
