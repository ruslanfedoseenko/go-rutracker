package Tests

import (
	"github.com/ruslanfedoseenko/gorutracker"
	"testing"
)

func TestGetForumNames(t *testing.T){
	client := gorutracker.NewClient()
	ids := []int{1, 2, 4, 5, 7}
	forumNames, err := client.GetForumNames(ids)
	if err != nil {
		t.Error(err)
		return
	}
	if len(forumNames) != len(ids){
		t.Error("Forum names reponse has incorrect length", forumNames)
		return
	}

	if len(forumNames) == 0 {
		t.Error("Empty forum names reponse ", forumNames)
		return
	}
}
