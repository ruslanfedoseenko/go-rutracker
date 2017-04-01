package Models

import (
	"github.com/ruslanfedoseenko/go-rutracker/Misc"
	"encoding/json"
	"errors"
)

type forumTopicData struct {
	TorrentStatusId 	int  	`json:"tor_status"`
	Seeders			int	`json:"seeders"`
	RegistrationUnixTime	int	`json:"reg_time"`
}

func (d *forumTopicData)  UnmarshalJSON(data []byte) (e error){
	//originally it is []int  [ "tor_status", "seeders", "reg_time"	]
	originalData := make([]int,3,3)
	e = json.Unmarshal(data, &originalData)
	if (e != nil) {
		return
	}

	if len(originalData) != 3 {
		return errors.New("Unable to unmarshal due to incorrect array length")
	}
	d = &forumTopicData{
		TorrentStatusId: 	originalData[0],
		Seeders: 		originalData[1],
		RegistrationUnixTime: 	originalData[2],
	}
	return
}

type ForumTopicDataResponse struct {
	Format 		 Misc.JSONString	`json:"format"`
	UpdateUnixTime 	 int			`json:"update_time"`
	UpdateTimeString string 		`json:"update_time_humn"`
	Result map[string]forumTopicData	`json:"result"`
}
