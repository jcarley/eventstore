package eventstore

type DomainEvent interface {
	Apply(es EventSourceAggregate)
}
