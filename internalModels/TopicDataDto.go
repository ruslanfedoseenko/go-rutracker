package internalModels

type TopicData struct {
	InfoHash, TopicTitle             string
	ForumId, PosterId, Size, Seeders int
	RegistrationUnixTime             int                        `json:"reg_time"`
	SeederLastSeenUnixTime           int                        `json:"seeder_last_seen"`
}

type TopicDataDto struct {
	Result map[string]TopicData
}
