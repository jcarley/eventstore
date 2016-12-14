package eventstore

type EventSourceAggregate interface {
	Version() int
	Changes() []DomainEvent
}

type Aggregate struct {
	changes []DomainEvent
	version int
}

func NewAggregate() *Aggregate {
	return &Aggregate{
		changes: make([]DomainEvent, 0, 5),
		version: 0,
	}
}

func (this *Aggregate) Version() int {
	return this.version
}

func (this *Aggregate) Changes() []DomainEvent {
	return this.changes
}

func (this *Aggregate) Apply(aggregate EventSourceAggregate, event DomainEvent) {
	// eventTypeName := reflect.TypeOf(change).Name()
	// methodName := fmt.Sprintf("When%s", eventTypeName)
	// this.Invoke(aggregate, methodName, change)

	// account := NewPayAsYouGoAccount()
	// domainEvents := []DomainEvents { ... }

	// for _, domainEvent := range domainEvents {
	// domainEvent.Apply(account)
	// }

	event.Apply(aggregate)
	this.version++
}

func (this *Aggregate) Causes(aggregate EventSourceAggregate, event DomainEvent) {
	this.changes = append(this.changes, event)
	this.Apply(aggregate, event)
}

// func (this *EventSourcedAggregate) Invoke(any interface{}, name string, args ...interface{}) {
// inputs := make([]reflect.Value, len(args))
// for i, _ := range args {
// inputs[i] = reflect.ValueOf(args[i])
// }
// reflect.ValueOf(any).MethodByName(name).Call(inputs)
// }
