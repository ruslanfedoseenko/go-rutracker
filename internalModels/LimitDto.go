package internalModels

type limitWrapper struct {
	Limit int                        `json:"limit"`
}

type LimitDto struct {
	Result limitWrapper              `json:"result"`
}
