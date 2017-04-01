package Models

type topicData struct {
	InfoHash, TopicTitle string
	ForumId, PosterId, Size, Seeders int
	RegistrationUnixTime int 			`json:"reg_time"`
	SeederLastSeenUnixTime int 			`json:"seeder_last_seen"`

}

type topicDataResponse struct {
	Result map[string]topicData
}
