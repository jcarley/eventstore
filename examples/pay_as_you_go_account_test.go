package examples

import (
	"fmt"
	"testing"
	"time"

	. "github.com/onsi/gomega"
)

func TestCallingApply(t *testing.T) {
	RegisterTestingT(t)

	// id := uuid.NewV4().String()
	// eventStore := providers.NewInMemoryEventStore()
	// repository := NewPayAsYouGoRepository(eventStore)

	account := NewPayAsYouGoAccount()
	fmt.Printf("%+v\n\n", account)

	account.IncreaseCreditLine(5)
	fmt.Printf("%+v\n\n", account)

	startTime, endTime := MakeCall()
	account.CallCompleted(startTime, endTime)
	fmt.Printf("%+v\n\n", account)

	account.IncreaseCreditLine(25)
	fmt.Printf("%+v\n\n", account)

	// repository.Save(account)

	fmt.Printf("%+v\n\n", account.Changes())
	fmt.Printf("Version: %d\n\n", account.Version())

	account2 := NewPayAsYouGoAccount()
	for _, change := range account.Changes() {
		account2.Apply(change)
	}
	fmt.Printf("%+v\n\n", account2)
	fmt.Printf("Version: %d\n\n", account2.Version())
}

func MakeCall() (time.Time, time.Time) {
	lengthOfCall, _ := time.ParseDuration("15m")

	startTime := time.Now()
	endTime := startTime.Add(lengthOfCall)

	return startTime, endTime
}
