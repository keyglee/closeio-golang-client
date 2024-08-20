package closeio_test

import (
	"testing"

	"github.com/keyglee/closeio-golang-client/closeio"
	"github.com/keyglee/closeio-golang-client/closeio/lead"
)

func TestCreateLead(t *testing.T) {

	newLead := lead.Lead{Name: "Name", Description: "Some Description"}

	client := closeio.CreateCloseClient("SomeApiKey")
	bytes, err := client.Create(newLead)

	if err != nil {
		t.Error(err)
	}
	t.Log(string(bytes))
}
