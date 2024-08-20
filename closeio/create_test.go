package closeio_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/keyglee/closeio-golang-client/closeio"
	"github.com/keyglee/closeio-golang-client/closeio/lead"
)

func TestCreateLead(t *testing.T) {
	apiKey := os.Getenv("CLOSE_IO_API_KEY")
	newLead := lead.Lead{Name: "Name", Description: "Some Description"}

	client := closeio.CreateCloseClient(apiKey)
	bytes, err := client.Create(newLead)

	if err != nil {
		t.Error(err)
	}
	t.Log(string(bytes))

	var response lead.LeadResponse

	err = json.Unmarshal(bytes, &response)

	if err != nil {
		t.Errorf("%+v", err)
	}

	t.Log(response)

	deleteLead := lead.Lead{ID: response.ID}
	bytes, err = client.Delete(deleteLead)

	if err != nil {
		t.Error(err)
	}

	t.Log(string(bytes))
}
