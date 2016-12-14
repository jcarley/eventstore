package examples

import (
	"fmt"
	"testing"
	"time"

	. "github.com/onsi/gomega"
	"github.com/twinj/uuid"
)

func TestCallingApply(t *testing.T) {
	RegisterTestingT(t)

	id := uuid.NewV4().String()

	eventStore := NewInMemoryEventStore()

	repository := NewPayAsYouGoRepository(eventStore)
	account, _ := repository.FindBy(id)
	fmt.Printf("%+v\n", account)

	account.IncreaseCreditLine(5)
	fmt.Printf("%+v\n", account)

	lengthOfCall, _ := time.ParseDuration("15m")

	startTime := time.Now()
	endTime := startTime.Add(lengthOfCall)
	account.CallCompleted(startTime, endTime)
	fmt.Printf("%+v\n", account)

	repository.Save(account)

	fmt.Printf("%+v\n", account.Changes())
	fmt.Printf("Version: %d\n", account.Version())
}
