package Tests

import (
	"testing"
	"github.com/ruslanfedoseenko/gorutracker/internalModels"
	"encoding/json"
)


func TestForumTopicDataResponseUnmarshalling(t *testing.T) {
	jsonString := "{\r\n  \"format\": {\r\n    \"topic_id\": [\r\n      \"tor_status\",\r\n      \"seeders\",\r\n      \"reg_time\"\r\n    ]\r\n  },\r\n  \"update_time\": 1491067831,\r\n  \"update_time_humn\": \"2017-04-01 20:30:31 MSK\",\r\n  \"total_size_bytes\": 20357180592027,\r\n  \"result\": {\r\n    \"111871\": [\r\n      2,\r\n      4,\r\n      1166721534\r\n    ],\r\n    \"228926\": [\r\n      2,\r\n      10,\r\n      1178191736\r\n    ]\r\n  }\r\n}"
	var forumTopicData internalModels.ForumTopicDataDto
	err := json.Unmarshal([]byte(jsonString), &forumTopicData)
	if err != nil {
		t.Error(err)
	}


}
