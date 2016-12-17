package examples

import (
	"fmt"
	"testing"
	"time"

	"github.com/jcarley/eventstore"
	"github.com/jcarley/eventstore/providers"
	. "github.com/onsi/gomega"
)

func TestCallingApply(t *testing.T) {
	RegisterTestingT(t)

	// id := uuid.NewV4().String()
	// eventStore := providers.NewInMemoryEventStore()

	eventStore := providers.NewPostgresEventStore()
	repository := NewPayAsYouGoRepository(eventStore)

	account := NewPayAsYouGoAccount()
	// account := repository.FindBy(id)
	fmt.Printf("%#v\n\n", account)

	account.IncreaseCreditLine(5)
	fmt.Printf("%#v\n\n", account)

	startTime, endTime := MakeCall(15)
	account.CallCompleted(startTime, endTime)
	fmt.Printf("%#v\n\n", account)

	account.IncreaseCreditLine(25)
	fmt.Printf("%#v\n\n", account)

	err := repository.Add(account)
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("%#v\n\n", account.Changes())
	fmt.Printf("Version: %d\n\n", account.Version())

	account2 := NewPayAsYouGoAccount()
	for _, change := range account.Changes() {
		account2.Apply(change)
	}
	fmt.Printf("%#v\n\n", account2)
	fmt.Printf("%#v\n\n", account2.Changes())
	fmt.Printf("Version: %d\n\n", account2.Version())
}

func TestProjectionApply(t *testing.T) {

	events := []eventstore.DomainEvent{
		NewPhoneCallCharged(15),
		NewPhoneCallCharged(10),
		NewPhoneCallCharged(25),
		NewPhoneCallCharged(65),
		CreditAdded{6},
		NewPhoneCallCharged(23),
		NewPhoneCallCharged(10),
		NewPhoneCallCharged(38),
	}

	projection := NewAverageCallDuration()
	for _, event := range events {
		projection.Apply(event)
	}

	fmt.Printf("\nThe total calls calculated: %d\n", projection.TotalCalls)
	fmt.Printf("Average length of call is %s minutes\n\n", projection.String())
}

func NewPhoneCallCharged(lengthOfCall int) PhoneCallCharged {
	startTime, endTime := MakeCall(lengthOfCall)
	costOfCall := float32(lengthOfCall) * 0.10

	return PhoneCallCharged{costOfCall, startTime.UTC().Format(time.RFC3339), endTime.UTC().Format(time.RFC3339)}
}

func MakeCall(lengthOfCall int) (time.Time, time.Time) {

	callDuration, _ := time.ParseDuration(fmt.Sprintf("%dm", lengthOfCall))

	startTime := time.Now()
	endTime := startTime.Add(callDuration)

	return startTime, endTime
}
