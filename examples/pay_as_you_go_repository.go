package examples

import (
	"fmt"
	"math"
	"reflect"

	"github.com/jcarley/eventstore"
)

type PayAsYouGoRepository struct {
	eventStore eventstore.EventStore
}

func NewPayAsYouGoRepository(eventStore eventstore.EventStore) *PayAsYouGoRepository {
	return &PayAsYouGoRepository{eventStore}
}

func (this *PayAsYouGoRepository) FindBy(id string) (*PayAsYouGoAccount, error) {

	streamName := this.StreamNameFor(id)
	fromEventNumber := 0
	toEventNumber := math.MaxInt32

	// snapShot, err := this.eventStore.GetLatestSnapshot(streamName)
	// if err != nil {
	// return nil, err
	// }

	// if snapShot != nil {
	// fromEventNumber = snapShot.Version + 1
	// }

	stream, err := this.eventStore.GetStream(streamName, fromEventNumber, toEventNumber)
	if err != nil {
		return nil, err
	}

	var account *PayAsYouGoAccount
	// if snapShot != nil {
	// account = NewPayAsYouGoAccountFromSnapshot(snapShot)
	// } else {
	account = NewPayAsYouGoAccount()
	// }
	account.ID = id
	for _, event := range stream {
		account.Apply(event)
	}

	return account, nil
}

func (this *PayAsYouGoRepository) Add(account *PayAsYouGoAccount) error {
	streamName := this.StreamNameFor(account.ID)
	return this.eventStore.CreateNewStream(streamName, account.Changes())
}

func (this *PayAsYouGoRepository) Save(account *PayAsYouGoAccount) error {
	streamName := this.StreamNameFor(account.ID)
	return this.eventStore.AppendEventsToStream(streamName, account.Changes(), 0)
}

func (this *PayAsYouGoRepository) StreamNameFor(id string) string {
	return fmt.Sprintf("%s-%s", reflect.TypeOf(&PayAsYouGoAccount{}).Elem().String(), id)
}
