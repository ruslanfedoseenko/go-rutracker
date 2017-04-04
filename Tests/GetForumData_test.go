package Tests

import (
	"testing"
	"github.com/ruslanfedoseenko/gorutracker"
)

func TestGetForumData(t *testing.T){
	client := gorutracker.NewClient()
	ids := []int{1, 2, 4, 5, 7}
	forumNames, err := client.GetForumData(ids)
	if err != nil {
		t.Error(err)
		return
	}
	if len(forumNames) != len(ids){
		t.Error("Forum data reponse has incorrect length", forumNames)
		return
	}

	if len(forumNames) == 0 {
		t.Error("Empty data names reponse ", forumNames)
		return
	}
}
