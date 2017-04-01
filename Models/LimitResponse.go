package Models

type limitWrapper struct {
	Limit int 			`json:"limit"`
}

type limitResponse struct{
	Result limitWrapper		`json:"result"`
}

