package Tests

import (
	"testing"
	"encoding/json"
	"github.com/ruslanfedoseenko/gorutracker/internalModels"
)

func TestForumNameUnmarshal(t *testing.T) {
	jsonString := "{\r\n  \"result\": {\r\n    \"1\": null,\r\n    \"7\": \"\u0417\u0430\u0440\u0443\u0431\u0435\u0436\u043D\u043E\u0435 \u043A\u0438\u043D\u043E\"\r\n  }\r\n}"
	var forumNameData internalModels.KeyValueDto
	err := json.Unmarshal([]byte(jsonString), &forumNameData)
	if err != nil {
		t.Error(err)
	}
}
