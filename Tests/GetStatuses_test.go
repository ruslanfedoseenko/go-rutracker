package Tests

import (
	"testing"
	"github.com/ruslanfedoseenko/gorutracker"
)

func TestGetStatuses (t *testing.T){
	client := gorutracker.NewClient()
	statuses, err := client.GetStatuses()
	if err != nil {
		t.Error(err)
	}

	if len(statuses ) == 0 {
		t.Error("Empty status reponse ", statuses)
	}
}
