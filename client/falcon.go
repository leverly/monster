package client

type Falcon7bInstructRequest struct {
	Model string               `json:"model"`
	Data  Falcon7bInstructData `json:"data"`
}

type Falcon7bInstructData struct {
	Prompt string  `json:"prompt"`
	TopK   int     `json:"top_k"`
	TopP   float64 `json:"top_p"`
}
