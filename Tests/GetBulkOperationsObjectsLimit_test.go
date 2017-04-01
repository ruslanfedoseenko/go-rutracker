package Tests

import (
	"testing"
	"github.com/ruslanfedoseenko/gorutracker"
)

func TestGetBulkOperationsObjectsLimit(t *testing.T){
	client := gorutracker.NewClient()
	limit, err := client.GetBulkOperationsObjectsLimit()
	if err != nil {
		t.Error(err)
	}

	if limit <=0 {
		t.Error("Limit", limit, "is negative or zero")
	}

}
