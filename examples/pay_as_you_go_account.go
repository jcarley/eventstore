package examples

import (
	"time"

	"github.com/jcarley/eventstore"
	"github.com/twinj/uuid"
)

type PayAsYouGoAccount struct {
	ID     string
	Amount float32
	*eventstore.Aggregate
}

func NewPayAsYouGoAccount() *PayAsYouGoAccount {
	return &PayAsYouGoAccount{
		ID:        uuid.NewV4().String(),
		Amount:    0,
		Aggregate: eventstore.NewAggregate(),
	}
}

func (this *PayAsYouGoAccount) IncreaseCreditLine(credit float32) {
	this.Aggregate.Causes(this, CreditAdded{credit})
}

func (this *PayAsYouGoAccount) CallCompleted(startTime, endTime time.Time) {
	elapsed := endTime.Sub(startTime)
	lengthOfCallInMinutes := elapsed.Minutes()
	costOfCall := float32(lengthOfCallInMinutes) * 0.10

	s := startTime.UTC().Format(time.RFC3339)
	e := endTime.UTC().Format(time.RFC3339)
	this.Aggregate.Causes(this, PhoneCallCharged{costOfCall, s, e})
}

func (this *PayAsYouGoAccount) Apply(change eventstore.DomainEvent) {
	this.Aggregate.Apply(this, change)
}

func (this *PayAsYouGoAccount) WhenCreditAdded(creditAdded CreditAdded) {
	this.Amount += creditAdded.Amount
}

func (this *PayAsYouGoAccount) WhenPhoneCallCharged(phoneCallCharged PhoneCallCharged) {
	this.Amount = this.Amount - phoneCallCharged.CostOfCall
}
