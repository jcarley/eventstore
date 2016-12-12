package eventstore

import (
	"fmt"
)

// type DomainEvent struct {
// }

type EventSourcedAggregate struct {
	changes []interface{}
	version int
}

func NewEventSourcedAggregate() *EventSourcedAggregate {
	return &EventSourcedAggregate{
		changes: make([]interface{}, 0, 5),
		version: 0,
	}
}

func (this *EventSourcedAggregate) Version() int {
	return this.version
}

func (this *EventSourcedAggregate) Changes() []interface{} {
	return this.changes
}

func (this *EventSourcedAggregate) Apply(aggregate interface{}, change interface{}) {
	fmt.Println("EventSourcedAggregate!Apply")
	// this.changes = append(this.changes, change)
	fmt.Printf("Aggregate: %T\n", aggregate)
	fmt.Printf("Event: %T\n", change)
	this.version++
}

// ************************************************

type CreditAdded struct {
}

type PhoneCallCharged struct {
}

type PayAsYouGoAccount struct {
	*EventSourcedAggregate
}

func NewPayAsYouGoAccount() *PayAsYouGoAccount {
	return &PayAsYouGoAccount{
		NewEventSourcedAggregate(),
	}
}

func (this *PayAsYouGoAccount) Apply(change interface{}) {
	fmt.Println("PayAsYouGoAccount!Apply")
	this.EventSourcedAggregate.Apply(this, change)
}

func (this *PayAsYouGoAccount) whenCreditAdded(creditAdded CreditAdded) {
	fmt.Println("PayAsYouGoAccount!whenCreditAdded")
}

func (this *PayAsYouGoAccount) whenPhoneCallCharged(phoneCallCharged PayAsYouGoAccount) {
	fmt.Println("PayAsYouGoAccount!whenPhoneCallCharged")
}
