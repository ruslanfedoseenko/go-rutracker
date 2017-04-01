package internalModels

type ForumData struct {
	ForumName string                        `json:"forum_name"`
	ParentId  int                           `json:"parent_id"`
}

type ForumDataDto struct {
	Result map[string]ForumData              `json:"result"`
}
