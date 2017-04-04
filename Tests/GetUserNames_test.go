package Tests

import (
	"testing"
	"github.com/ruslanfedoseenko/gorutracker"
)

func TestGetUserNames(t *testing.T){
	client := gorutracker.NewClient()
	ids := []int{1, 2, 4, 5, 7}
	userNames, err := client.GetUserName(ids)
	if err != nil {
		t.Error(err)
		return
	}
	if len(userNames) != len(ids){
		t.Error("Forum names reponse has incorrect length", userNames)
		return
	}

	if len(userNames) == 0 {
		t.Error("Empty forum names reponse ", userNames)
		return
	}
}
