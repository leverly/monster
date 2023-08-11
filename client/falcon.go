package client

type FalconTextGenerationRequest struct {
	Model string                    `json:"model"`
	Param FalconTextGenerationParam `json:"data"`
}

type FalconTextGenerationParam struct {
	Prompt string  `json:"prompt"`
	TopK   int     `json:"top_k"`
	TopP   float64 `json:"top_p"`
}
