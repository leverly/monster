package client

type TextGenerationRequest struct {
	Model string             `json:"model"`
	Data  TextGenerationData `json:"data"`
}

type TextGenerationData struct {
	Prompt string  `json:"prompt"`
	TopK   int     `json:"top_k"`
	TopP   float64 `json:"top_p"`
}
