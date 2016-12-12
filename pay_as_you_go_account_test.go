package eventstore

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"
	"github.com/twinj/uuid"
)

func TestCallingApply(t *testing.T) {
	RegisterTestingT(t)

	id := uuid.NewV4().String()

	repository := NewPayAsYouGoRepository(nil)
	account, _ := repository.FindBy(id)
	account.TopUp(5)
	repository.Save(account)

	fmt.Println(len(account.Changes()))
	fmt.Printf("%+v\n", account.Changes())
	fmt.Printf("Version: %d\n", account.Version())
}
