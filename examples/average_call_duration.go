package examples

import (
	"fmt"
	"time"

	"github.com/jcarley/eventstore"
)

type AverageCallDuration struct {
	TotalCalls        int
	CallLengthTotal   float32
	AverageCallLength float32
	*eventstore.Aggregate
}

func NewAverageCallDuration() *AverageCallDuration {
	return &AverageCallDuration{
		TotalCalls:        0,
		CallLengthTotal:   0,
		AverageCallLength: 0,
		Aggregate:         eventstore.NewAggregate(),
	}
}

func (this *AverageCallDuration) Apply(event eventstore.DomainEvent) {
	this.Aggregate.Apply(this, event)
}

func (this *AverageCallDuration) WhenPhoneCallCharged(phoneCallCharged PhoneCallCharged) {
	this.TotalCalls++

	startTime, _ := time.Parse(time.RFC3339, phoneCallCharged.StartTime)
	endTime, _ := time.Parse(time.RFC3339, phoneCallCharged.EndTime)

	elapsed := endTime.Sub(startTime)
	lengthOfCallInMinutes := elapsed.Minutes()
	this.CallLengthTotal += float32(lengthOfCallInMinutes)

	this.AverageCallLength = this.CallLengthTotal / float32(this.TotalCalls)
}

func (this *AverageCallDuration) String() string {
	return fmt.Sprintf("%.2f", this.AverageCallLength)
}
