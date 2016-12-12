package eventstore

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
)

func TestCallingApply(t *testing.T) {
	RegisterTestingT(t)

	account := NewPayAsYouGoAccount()

	change := CreditAdded{}
	account.Apply(change)

	fmt.Println(len(account.Changes()))
	fmt.Println(account.Changes())
	fmt.Printf("Version: %d\n", account.Version())

}
