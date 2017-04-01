package Models

type forumData struct {
	ForumName string			`json:"forum_name"`
	ParentId int				`json:"parent_id"`
}

type forumDataResponse struct {
	Result map[string]forumData		`json:"result"`
}
