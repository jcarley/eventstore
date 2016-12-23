package examples

import "github.com/jcarley/eventstore"

func init() {
	eventstore.RegisterType((*CreditAdded)(nil))
}

type CreditAddedTarget interface {
	WhenCreditAdded(CreditAdded)
}

type CreditAdded struct {
	Amount float32 `json:"amount,omitempty"`
}

func (this CreditAdded) Apply(es eventstore.EventSourceAggregate) {
	if target, ok := es.(CreditAddedTarget); ok {
		target.WhenCreditAdded(this)
	}
}
