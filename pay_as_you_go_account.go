package eventstore

import (
	"fmt"

	"github.com/twinj/uuid"
)

type ResizeMessage struct {
	Source      string `json:"source,omitempty"`
	Destination string `json:"destination,omitempty"`
	Size        string `json:"size,omitempty"`
}

type CreditAdded struct {
	Amount int `json:"amount,omitempty"`
}

type PhoneCallCharged struct {
	Minutes int `json:"minutes,omitempty"`
}

type PayAsYouGoAccount struct {
	ID string
	*EventSourcedAggregate
}

func NewPayAsYouGoAccount() *PayAsYouGoAccount {
	return &PayAsYouGoAccount{
		uuid.NewV4().String(),
		NewEventSourcedAggregate(),
	}
}

func (this *PayAsYouGoAccount) Apply(change interface{}) {
	this.EventSourcedAggregate.Apply(this, change)
}

func (this *PayAsYouGoAccount) TopUp(credit int) {
	this.EventSourcedAggregate.Causes(this, CreditAdded{credit})
}

func (this *PayAsYouGoAccount) WhenCreditAdded(creditAdded CreditAdded) {
	fmt.Println("PayAsYouGoAccount!WhenCreditAdded")
}

func (this *PayAsYouGoAccount) WhenPhoneCallCharged(phoneCallCharged PhoneCallCharged) {
	fmt.Println("PayAsYouGoAccount!WhenPhoneCallCharged")
}
