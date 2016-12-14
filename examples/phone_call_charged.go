package examples

import "github.com/jcarley/eventstore"

type PhoneCallChargedTarget interface {
	WhenPhoneCallCharged(PhoneCallCharged)
}

type PhoneCallCharged struct {
	CostOfCall float32 `json:"costOfCall,omitempty"`
	StartTime  string
	EndTime    string
}

func (this PhoneCallCharged) Apply(es eventstore.EventSourceAggregate) {
	if target, ok := es.(PhoneCallChargedTarget); ok {
		target.WhenPhoneCallCharged(this)
	}
}
