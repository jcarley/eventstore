package examples

import (
	"fmt"
	"testing"
	"time"

	"github.com/jcarley/eventstore"
	"github.com/jcarley/eventstore/providers"
	. "github.com/onsi/gomega"
)

func TestAddingNewEventStream(t *testing.T) {
	RegisterTestingT(t)

	eventStore := providers.NewPostgresEventStore()
	repository := NewPayAsYouGoRepository(eventStore)

	account := NewPayAsYouGoAccount()
	account.IncreaseCreditLine(5)
	startTime, endTime := MakeCall(15)
	account.CallCompleted(startTime, endTime)
	account.IncreaseCreditLine(25)

	err := repository.Add(account)
	if err != nil {
		t.Error(err)
	}

	// Assert
	Expect(len(account.Changes())).To(Equal(0), "Changes should be cleared")
	Expect(account.Version()).To(Equal(3), "Incorrect version number")
}

func TestSavingToAnEventStream(t *testing.T) {
	RegisterTestingT(t)

	// Arrange
	eventStore := providers.NewPostgresEventStore()
	repository := NewPayAsYouGoRepository(eventStore)
	account := NewPayAsYouGoAccount()
	account.IncreaseCreditLine(5)
	startTime, endTime := MakeCall(15)
	account.CallCompleted(startTime, endTime)
	account.IncreaseCreditLine(25)
	err := repository.Add(account)
	if err != nil {
		t.Error(err)
	}

	// Action
	account.IncreaseCreditLine(5)
	startTime, endTime = MakeCall(35)
	account.CallCompleted(startTime, endTime)

	Expect(len(account.Changes())).To(Equal(2), "Changes should be cleared")

	err = repository.Save(account)
	if err != nil {
		t.Error(err)
	}

	// Assert
	Expect(len(account.Changes())).To(Equal(0), "Changes should be cleared")
	Expect(account.Version()).To(Equal(5), "Incorrect version number")
}

func TestProjectionApply(t *testing.T) {
	RegisterTestingT(t)

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

	actualCallLength := fmt.Sprintf("%.2f", projection.AverageCallLength)
	expectedCallLength := "26.57"

	Expect(projection.TotalCalls).To(Equal(7))
	Expect(actualCallLength).To(Equal(expectedCallLength))
}

func TestGetStream(t *testing.T) {
	RegisterTestingT(t)

	eventStore := providers.NewPostgresEventStore()
	repository := NewPayAsYouGoRepository(eventStore)

	account := NewPayAsYouGoAccount()
	account.IncreaseCreditLine(5)
	startTime, endTime := MakeCall(15)
	account.CallCompleted(startTime, endTime)
	account.IncreaseCreditLine(25)

	err := repository.Add(account)
	if err != nil {
		t.Error(err)
	}

	id := account.ID
	fmt.Println(id)

	account2, err := repository.FindBy(id)

	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
	}

	fmt.Println("====================================")
	fmt.Printf("%#v\n", account2)
	fmt.Println("====================================")
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
